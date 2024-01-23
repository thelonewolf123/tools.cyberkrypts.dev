package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kkdai/youtube/v2"
)

type Format struct {
	QualityLabel string
	URL          string
	MimeType     string
}

type VideoResponse struct {
	Title   string
	Author  string
	Formats []Format
}

func getVideoInfo(videoURL string) (VideoResponse, error) {
	youtube := youtube.Client{}
	video, err := youtube.GetVideo(videoURL)

	if err != nil {
		return VideoResponse{}, err
	}

	formats := []Format{}

	for _, format := range video.Formats {
		formats = append(formats, Format{QualityLabel: format.QualityLabel, URL: format.URL, MimeType: format.MimeType})
	}

	videoResponse := VideoResponse{
		Title:   video.Title,
		Author:  video.Author,
		Formats: formats,
	}

	return videoResponse, nil
}

// https://www.youtube.com/watch?v=HAo_YVzRelk
func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{})
	})

	router.GET("/video", func(c *gin.Context) {
		videoURL := c.Query("url")

		if videoURL == "" {
			c.JSON(400, gin.H{"error": "url is required"})
			return
		}

		videoResponse, err := getVideoInfo(videoURL)

		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, videoResponse)
	})

	router.Run(":8080")
}

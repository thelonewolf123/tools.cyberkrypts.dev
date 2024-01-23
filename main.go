package main

import (
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kkdai/youtube/v2"
)

type Format struct {
	QualityLabel string
	URL          string
	FileName     string
	FileType     string
}

type VideoResponse struct {
	Title   string
	Author  string
	Formats []Format
}

func getFileName(title, mimeType string) (string, string) {
	mimeType = strings.Split(mimeType, ";")[0]
	switch mimeType {
	case "video/mp4":
		return title + ".mp4", "mp4"
	case "video/webm":
		return title + ".webm", "webm"
	case "video/3gpp":
		return title + ".3gp", "3gp"
	case "video/x-flv":
		return title + ".flv", "flv"
	case "video/quicktime":
		return title + ".mov", "mov"
	case "audio/mp4":
		return title + ".m4a", "m4a"
	case "audio/webm":
		return title + ".webm", "webm"
	case "audio/3gpp":
		return title + ".3gp", "3gp"
	}

	return title + ".mp4", "mp4"
}

func getVideoInfo(videoURL string) (VideoResponse, error) {
	youtube := youtube.Client{}
	video, err := youtube.GetVideo(videoURL)

	if err != nil {
		return VideoResponse{}, err
	}

	formats := []Format{}

	for _, format := range video.Formats {
		fileName, fileType := getFileName(video.Title, format.MimeType)
		qualityLabel := format.QualityLabel
		if qualityLabel == "" {
			qualityLabel = "Audio"
		}
		formats = append(formats, Format{QualityLabel: qualityLabel, URL: format.URL, FileName: fileName, FileType: fileType})
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

		c.HTML(200, "video.html", gin.H{
			"title":   videoResponse.Title,
			"author":  videoResponse.Author,
			"formats": videoResponse.Formats,
		})
	})

	if os.Getenv("PORT") != "" {
		router.Run(":" + os.Getenv("PORT"))
		return
	}

	router.Run(":8080")
}

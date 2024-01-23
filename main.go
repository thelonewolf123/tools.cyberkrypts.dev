package main

import (
	"fmt"
	"log"
	"os"

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

// https://www.youtube.com/watch?v=HAo_YVzRelk
func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <YouTube Video URL>")
		return
	}

	videoURL := os.Args[1]

	youtube := youtube.Client{}
	video, err := youtube.GetVideo(videoURL)

	if err != nil {
		log.Fatalf("Error getting video: %v", err)
	}

	formats := []Format{}

	for _, format := range video.Formats {
		fmt.Printf("Format: %s\n", format.QualityLabel)
		fmt.Printf("URL: %s\n", format.URL)
		fmt.Printf("Mime Type: %s\n", format.MimeType)

		formats = append(formats, Format{QualityLabel: format.QualityLabel, URL: format.URL, MimeType: format.MimeType})
	}

	videoResponse := VideoResponse{
		Title:   video.Title,
		Formats: formats,
		Author:  video.Author,
	}

	fmt.Printf("Video Title: %s\n", videoResponse)
}

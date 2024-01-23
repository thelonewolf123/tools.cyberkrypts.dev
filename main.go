package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kkdai/youtube/v2"
)

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

	streamUrl, err := youtube.GetStreamURL(video, &video.Formats[0])

	if err != nil {
		log.Fatalf("Error getting stream: %v", err)
	}

	fmt.Printf("Video Title: %s\n", video.Title)
	fmt.Printf("Video Author: %s\n", streamUrl)
}

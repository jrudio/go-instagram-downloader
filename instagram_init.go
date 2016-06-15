package main

import (
	"flag"
	"fmt"
	"os"
)

func initFlags() {
	flag.StringVar(&igURL, "url", "", "Ex: https://www.instagram.com/badgalriri")
	flag.StringVar(&igUsername, "username", "", "Ex: badgalriri")
	flag.StringVar(&outputPath, "o", "./", "Ex: /home/bob/downloads or ../pictures")

	flag.IntVar(&downloadCount, "count", 0, "Ex: Download the first n media; 0 == max")

	flag.BoolVar(&grabImages, "images", false, "Grab standard resolution images")
	flag.BoolVar(&grabLowResImages, "images-low", false, "Grab low resolution images")
	flag.BoolVar(&grabThumbnails, "thumb", false, "Grab thumbnails")
	flag.BoolVar(&grabVideos, "videos", false, "Grab standard resolution videos")
	flag.BoolVar(&grabLowBandwithVideos, "videos-lowband", false, "Grab low bandwith videos")
	flag.BoolVar(&grabLowResVideos, "videos-lowres", false, "Grab low resolution videos")
	flag.Parse()

	if igURL == "" && igUsername == "" {
		fmt.Println("An instagram url or username is required")
		flag.Usage()
		os.Exit(1)
	}
}

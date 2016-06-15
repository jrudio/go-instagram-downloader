package main

import (
	"fmt"
	"sync"
)

func main() {
	initFlags()

	request, requestErr := newRequest(igUsername, igURL)

	if requestErr != nil {
		fmt.Printf("New Request: %s\n", requestErr.Error())
		return
	}

	igPage, err := get(request.fullURL)

	if err != nil {
		fmt.Printf("Fetch: %s\n", err.Error())
		return
	}

	mediaType := "images"
	var mediaList []string

	// Order of preference
	if grabImages {
		fmt.Println("Grabbing standard resolution images")
		mediaList = getStandardResImages(igPage)
	} else if grabLowResImages {
		fmt.Println("Grabbing low resolution images")
		mediaList = getLowResImages(igPage)
	} else if grabThumbnails {
		fmt.Println("Grabbing thumbnails")
		mediaList = getThumbnails(igPage)
	} else if grabVideos {
		fmt.Println("Grabbing videos")
		mediaList = getStandardResVideos(igPage)
		mediaType = "videos"
	} else if grabLowBandwithVideos {
		fmt.Println("Grabbing low bandwith videos")
		mediaList = getLowBandwithVideos(igPage)
		mediaType = "videos"
	} else if grabLowResVideos {
		fmt.Println("Grabbing videos")
		mediaList = getLowResVideos(igPage)
		mediaType = "videos"
	} else {
		fmt.Println("Grabbing standard resolution images")
		mediaList = getStandardResImages(igPage)
	}

	if outputPath == "-" {
		count := 0
		for _, m := range mediaList {
			if len(m) == 0 {
				continue
			}

			fmt.Println(m)

			count++

			if downloadCount != 0 && downloadCount == count {
				break
			}
		}
		return
	}

	// Download list
	cb := saveImage

	if mediaType == "videos" {
		cb = saveVideo
	}

	var wg sync.WaitGroup

	count := 0

	for ii, m := range mediaList {
		if len(m) == 0 {
			continue
		}

		wg.Add(1)
		count++
		go func(query string, currentIndex, count int) {
			defer wg.Done()

			if err = cb(query, fmt.Sprintf("%s-%d", request.username, currentIndex)); err != nil {
				fmt.Printf("Save media: %s\n", err.Error())
			}
		}(m, ii, count)

		if downloadCount != 0 && downloadCount == count {
			break
		}
	}

	wg.Wait()
	fmt.Printf("Downloaded all public media for %s\n", request.username)
}

// Images

func getThumbnails(ig instagramPage) []string {
	itemCount := len(ig.Items)
	thumbnails := make([]string, itemCount)

	for ii, item := range ig.Items {
		url := item.Images.Thumbnail.URL

		thumbnails[ii] = url
	}

	return thumbnails
}

func getLowResImages(ig instagramPage) []string {
	itemCount := len(ig.Items)
	images := make([]string, itemCount)

	for ii, item := range ig.Items {
		url := item.Images.LowResolution.URL

		images[ii] = url
	}

	return images
}

func getStandardResImages(ig instagramPage) []string {
	itemCount := len(ig.Items)
	images := make([]string, itemCount)

	for ii, item := range ig.Items {
		url := item.Images.StandardResolution.URL

		images[ii] = url
	}

	return images
}

// Videos

func getLowBandwithVideos(ig instagramPage) []string {
	itemCount := len(ig.Items)
	videos := make([]string, itemCount)

	for ii, item := range ig.Items {
		url := item.Videos.LowBandwidth.URL

		videos[ii] = url
	}

	return videos
}

func getLowResVideos(ig instagramPage) []string {
	itemCount := len(ig.Items)
	videos := make([]string, itemCount)

	for ii, item := range ig.Items {
		url := item.Videos.LowResolution.URL

		videos[ii] = url
	}

	return videos
}

func getStandardResVideos(ig instagramPage) []string {
	itemCount := len(ig.Items)
	videos := make([]string, itemCount)

	for ii, item := range ig.Items {
		url := item.Videos.StandardResolution.URL

		videos[ii] = url
	}

	return videos
}

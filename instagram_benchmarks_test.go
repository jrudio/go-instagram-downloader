package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func getTestIGPage(data string) (instagramPage, error) {
	var igPage instagramPage
	if err := json.Unmarshal([]byte(data), &igPage); err != nil {
		return instagramPage{}, err
	}

	return igPage, nil
}

func BenchmarkFindThumbnails(b *testing.B) {
	igPage, err := getTestIGPage(testData)

	if err != nil {
		fmt.Printf("Test ig page: %s\n", err.Error())
		return
	}

	b.ResetTimer()

	for ii := 0; ii < b.N; ii++ {
		getThumbnails(igPage)
	}
}

func BenchmarkFindLowResImages(b *testing.B) {
	igPage, err := getTestIGPage(testData)

	if err != nil {
		fmt.Printf("Test ig page: %s\n", err.Error())
		return
	}

	b.ResetTimer()

	for ii := 0; ii < b.N; ii++ {
		getLowResImages(igPage)
	}
}

func BenchmarkFindStandardResImages(b *testing.B) {
	igPage, err := getTestIGPage(testData)

	if err != nil {
		fmt.Printf("Test ig page: %s\n", err.Error())
		return
	}

	b.ResetTimer()

	for ii := 0; ii < b.N; ii++ {
		getStandardResImages(igPage)
	}
}

func BenchmarkFindStandardResVideos(b *testing.B) {
	igPage, err := getTestIGPage(testData)

	if err != nil {
		fmt.Printf("Test ig page: %s\n", err.Error())
		return
	}

	b.ResetTimer()

	for ii := 0; ii < b.N; ii++ {
		getStandardResVideos(igPage)
	}
}

func BenchmarkFindLowResVideos(b *testing.B) {
	igPage, err := getTestIGPage(testData)

	if err != nil {
		fmt.Printf("Test ig page: %s\n", err.Error())
		return
	}

	b.ResetTimer()

	for ii := 0; ii < b.N; ii++ {
		getLowResVideos(igPage)
	}
}

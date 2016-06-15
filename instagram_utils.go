package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func newRequest(username, fullURL string) (*igRequest, error) {
	_request := &igRequest{username: username, fullURL: fullURL}

	if fullURL == "" && username != "" {
		_request.fullURL = fmt.Sprintf("https://www.instagram.com/%s/media", username)
	} else if username == "" && fullURL != "" {
		var nameErr error

		_request.username, nameErr = getUsernameFromURL(fullURL)

		if nameErr != nil {
			return _request, nameErr
		}
	}

	return _request, nil
}

func get(query string) (instagramPage, error) {
	resp, respErr := http.Get(query)

	if respErr != nil {
		return instagramPage{}, fmt.Errorf("Response: %s\n", respErr.Error())
	}

	defer resp.Body.Close()

	var igPage instagramPage

	if err := json.NewDecoder(resp.Body).Decode(&igPage); err != nil {
		return instagramPage{}, fmt.Errorf("JSON: %s\n", err.Error())
	}

	return igPage, nil
}

func saveImage(query, fileNamePrefix string) error {
	resp, respErr := http.Get(query)

	if respErr != nil {
		return fmt.Errorf("Response: %s\n", respErr.Error())
	}

	defer resp.Body.Close()

	extensionName := "png"
	contentType := resp.Header.Get("content-type")

	if len(contentType) >= 6 {
		contentType = contentType[6:]
		extensionName = contentType
	}

	if outPathCount := len(outputPath); outPathCount > 2 {
		if end := outputPath[outPathCount-1:]; end != "/" {
			outputPath += "/"
		}
	}

	file, err := os.OpenFile(fmt.Sprintf("%s%s.%s", outputPath, fileNamePrefix, extensionName), os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)

	if err != nil {
		return fmt.Errorf("Create file: %s\n", err.Error())
	}

	defer file.Close()

	_, writeErr := io.Copy(file, resp.Body)

	if writeErr != nil {
		return fmt.Errorf("Image write: %s\n", err.Error())
	}

	return nil
}

func saveVideo(query, fileNamePrefix string) error {
	resp, respErr := http.Get(query)

	if respErr != nil {
		return fmt.Errorf("Response: %s\n", respErr.Error())
	}

	defer resp.Body.Close()

	file, err := os.OpenFile(fmt.Sprintf("%s%s.mp4", outputPath, fileNamePrefix), os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)

	if err != nil {
		return fmt.Errorf("Create file: %s\n", err.Error())
	}

	defer file.Close()

	_, writeErr := io.Copy(file, resp.Body)

	if writeErr != nil {
		return fmt.Errorf("Video write: %s\n", err.Error())
	}

	return nil
}

func getUsernameFromURL(_url string) (string, error) {
	urlCount := len(_url)

	// https://www.instagram.com/ == 26 chars
	if urlCount <= 26 {
		return "", errors.New("URL is not long enough to determine username")
	}

	_url = _url[26:]

	for ii := 0; ii < urlCount; ii++ {
		if _url[ii:ii+1] == "/" {
			return _url[:ii], nil
		}
	}

	return _url, errors.New("Username not found")
}

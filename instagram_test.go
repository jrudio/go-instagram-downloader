package main

import "testing"

// TestNewRequest edge cases

func TestNewRequest(t *testing.T) {
	username := "badgalriri"
	fullURL := "https://www.instagram.com/badgalriri/media"

	request, err := newRequest(username, fullURL)

	if err != nil {
		t.Error(err.Error())
		return
	}

	if request.username != username {
		t.Errorf("Username does not match initialized value\n\tExpected: %s; Got: %s\n", username, request.username)
	}

	if request.fullURL != fullURL {
		t.Errorf("URL does not match initialized value\n\tExpected: %s; Got: %s\n", fullURL, request.fullURL)
	}
}

func TestNewRequestEmptyUsername(t *testing.T) {
	username := ""
	expectedUsername := "badgalriri"
	fullURL := "https://www.instagram.com/badgalriri/media"

	request, err := newRequest(username, fullURL)

	if err != nil {
		t.Error(err.Error())
		return
	}

	if request.username != expectedUsername {
		t.Errorf("Expected: %s; Got: %s\n", expectedUsername, request.username)
	}
}

func TestNewRequestEmptyURL(t *testing.T) {
	username := "badgalriri"
	expectedURL := "https://www.instagram.com/badgalriri/media"
	fullURL := ""

	request, err := newRequest(username, fullURL)

	if err != nil {
		t.Error(err.Error())
		return
	}

	if request.fullURL != expectedURL {
		t.Errorf("Expected: %s; Got: %s\n", expectedURL, request.fullURL)
	}
}

func TestNewRequestAllEmpty(t *testing.T) {
	username := ""
	fullURL := ""

	_, err := newRequest(username, fullURL)

	if err != nil {
		t.Error(err.Error())
	}
}

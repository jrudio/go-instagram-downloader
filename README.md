# Instagram downloader
 Instagram image and video downloader written in Go. Inspiration for this project came from browsing [this /r/forhire/ thread](https://www.reddit.com/r/forhire/comments/4nyf30/hiring_instagram_bot_sort_content_by_likes_views/d49711i) and found out you can get cool information in json format.
 
 
## Usage

```bash
$ ig -help

Usage of ig:
  -count int
    	Ex: Download the first n media; 0 == max
  -images
    	Grab standard resolution images
  -images-low
    	Grab low resolution images
  -o string
    	Ex: /home/bob/downloads or ../pictures (default "./")
  -thumb
    	Grab thumbnails
  -url string
    	Ex: https://www.instagram.com/badgalriri
  -username string
    	Ex: badgalriri
  -videos
    	Grab standard resolution videos
  -videos-lowband
    	Grab low bandwith videos
  -videos-lowres
    	Grab low resolution videos
```

1. Using `ig -username badgalriri` will download the 20 most-recent media on the user's public profile

2. Use `ig -username badgalriri -count 2` to limit the items to download

3. Use `ig -username badgalriri -o -` to output the media links to stdout

4. That's kind of it
 
## How to compile this program

1. Install [Go](https://golang.org/dl/)

2. Next `go get -u github.com/jrudio/go-instagram-downloader`

##### Compiling from source note: Change the name of the executable to something short like `ig`

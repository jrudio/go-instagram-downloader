package main

type igRequest struct {
	username string
	fullURL  string
}

type instagramPage struct {
	Items []struct {
		AltMediaURL       string `json:"alt_media_url"`
		CanDeleteComments bool   `json:"can_delete_comments"`
		CanViewComments   bool   `json:"can_view_comments"`
		Caption           struct {
			CreatedTime string `json:"created_time"`
			From        struct {
				FullName       string `json:"full_name"`
				ID             string `json:"id"`
				ProfilePicture string `json:"profile_picture"`
				Username       string `json:"username"`
			} `json:"from"`
			ID   string `json:"id"`
			Text string `json:"text"`
		} `json:"caption"`
		Code     string `json:"code"`
		Comments struct {
			Count int `json:"count"`
			Data  []struct {
				CreatedTime string `json:"created_time"`
				From        struct {
					FullName       string `json:"full_name"`
					ID             string `json:"id"`
					ProfilePicture string `json:"profile_picture"`
					Username       string `json:"username"`
				} `json:"from"`
				ID   string `json:"id"`
				Text string `json:"text"`
			} `json:"data"`
		} `json:"comments"`
		CreatedTime string `json:"created_time"`
		ID          string `json:"id"`
		Images      struct {
			LowResolution struct {
				Height int    `json:"height"`
				URL    string `json:"url"`
				Width  int    `json:"width"`
			} `json:"low_resolution"`
			StandardResolution struct {
				Height int    `json:"height"`
				URL    string `json:"url"`
				Width  int    `json:"width"`
			} `json:"standard_resolution"`
			Thumbnail struct {
				Height int    `json:"height"`
				URL    string `json:"url"`
				Width  int    `json:"width"`
			} `json:"thumbnail"`
		} `json:"images"`
		Likes struct {
			Count int `json:"count"`
			Data  []struct {
				FullName       string `json:"full_name"`
				ID             string `json:"id"`
				ProfilePicture string `json:"profile_picture"`
				Username       string `json:"username"`
			} `json:"data"`
		} `json:"likes"`
		Link     string      `json:"link"`
		Location interface{} `json:"location"`
		Type     string      `json:"type"`
		User     struct {
			FullName       string `json:"full_name"`
			ID             string `json:"id"`
			ProfilePicture string `json:"profile_picture"`
			Username       string `json:"username"`
		} `json:"user"`
		UserHasLiked bool `json:"user_has_liked"`
		VideoViews   int  `json:"video_views"`
		Videos       struct {
			LowBandwidth struct {
				Height int    `json:"height"`
				URL    string `json:"url"`
				Width  int    `json:"width"`
			} `json:"low_bandwidth"`
			LowResolution struct {
				Height int    `json:"height"`
				URL    string `json:"url"`
				Width  int    `json:"width"`
			} `json:"low_resolution"`
			StandardResolution struct {
				Height int    `json:"height"`
				URL    string `json:"url"`
				Width  int    `json:"width"`
			} `json:"standard_resolution"`
		} `json:"videos"`
	} `json:"items"`
	MoreAvailable bool   `json:"more_available"`
	Status        string `json:"status"`
}

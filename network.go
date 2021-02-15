package main

import (
	"log"
	"net/http"
)

// getUrlResponse get and check image url with two resolutions.
// If "/maxresdefault.jpg" have a bad response or doesn't exist
// getURLResponse try to get url with lowest or only resolution "/hqdefault.jpg".
func (t *Thumbnail) getURLResponse() *http.Response {

	// two possible resolutions
	const (
		vi     = "https://i.ytimg.com/vi/"
		resMax = "/maxresdefault.jpg"
		resHQ  = "/hqdefault.jpg"
	)

	resp, err := http.Get(vi + t.VideoID + resMax)

	if err != nil || resp.StatusCode != 200 {

		resp, err = http.Get(vi + t.VideoID + resHQ)

		if err != nil || resp.StatusCode != 200 {
			log.Printf("Response Status Code: %v\n", resp.StatusCode)
			return nil
		}
	}
	return resp
}

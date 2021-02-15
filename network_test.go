package main

import "testing"

func TestGetURLResponse(t *testing.T) {

	idList := []string{"l7xtyI2_MAQ", "G1IbRujko-A", "ZnhquCll3uQ"}

	tmb := &Thumbnail{}

	for _, id := range idList {
		tmb.VideoID = id
		if tmb.getURLResponse() == nil {
			t.Error("Must be non nil return!")
		}
	}
}

func TestGetWrongURLResponse(t *testing.T) {

	idList := []string{"N2wdg7Kad5F", "65ADa3Lsd9Q", "ZnhqWQ42I6w"}

	tmb := &Thumbnail{}

	for _, id := range idList {
		tmb.VideoID = id
		if tmb.getURLResponse() != nil {
			t.Error("Must be nil return!")
		}
	}
}

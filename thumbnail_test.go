package main

import (
	"fmt"
	"path/filepath"
	"testing"
)

func TestFindVideoID(t *testing.T) {

	// Allow next url types:
	// https://www.youtube.com/watch?v=N2wJQSBx5i4												standart
	// https://www.youtube.com/watch?v=65AB2pMCj4I&index=3&list=LLdt97678HxmYdM0DyZ847Uw		playlist
	// https://youtu.be/ZnhquCll3uQ																short url
	// https://www.youtube.com/embed/5eNieKeLBLE												embed url
	// https://youtu.be/6k1oE2y7NIo?t=31														timecode url

	tmb := &Thumbnail{}

	urlList := []string{
		"https://www.youtube.com/watch?v=N2wJQSBx5i4",
		"https://www.youtube.com/watch?v=65AB2pMCj4I&index=3&list=LLdt97678HxmYdM0DyZ847Uw",
		"https://youtu.be/ZnhquCll3uQ",
		"https://www.youtube.com/embed/5eNieKeLBLE",
		"https://youtu.be/6k1oE2y7NIo?t=31",
	}

	for _, url := range urlList {

		if tmb.findVideoID(url) != nil {
			fmt.Println("Corrupted Video URL:", tmb.VideoID)
			t.Errorf("Must be nil return! Link %s is invalid!", url)
		}
	}
}

func TestFindVideoWrongID(t *testing.T) {

	tmb := &Thumbnail{}

	urlList := []string{
		"https://www.youtube.com/watch?v=N2wJ" + "?" + "SBx5i4",
		"https://www.youtube.com/watch?v=65AB" + "&" + "pMCj4I&index=3&list=LLdt97678HxmYdM0DyZ847Uw",
		"https://youtu.be/Znh+" + "<" + "uC" + "%" + "l3uQ",
		"https://www.youtube.com/embed/5e" + "&" + "iKe" + "/" + "BLE",
	}

	for _, url := range urlList {
		if tmb.findVideoID(url) == nil {
			t.Errorf("Must be non nil return! Link ID %s is invalid!", url)
		}
	}
}

func TestFindVideoWrongIDLength(t *testing.T) {

	tmb := &Thumbnail{}

	urlList := []string{
		"https://www.youtube.com/watch?v=N2wJQ",
		"https://www.youtube.com/watch?v=65AB2p4I&index=3&list=LLdt97678HxmYdM0DyZ847Uw",
		"https://youtu.be/ZnhquCl",
		"https://www.youtube.com/embed/5eNi",
	}

	for _, url := range urlList {
		if tmb.findVideoID(url) == nil {
			t.Errorf("Must be non nil return! Link ID length %s is invalid!", url)
		}
	}
}

func TestWalkFunc(t *testing.T) {

	tmb := new(Thumbnail)
	thumbnailsDir := "./test_data"

	err := filepath.Walk(thumbnailsDir, tmb.walkFunc)
	if err != nil {
		t.Error("Walk Func Test Failed. Must retrun nil!")
	}
}

func TestSetNameDigit(t *testing.T) {

	namesList := []string{"thumbnail_0.jpg", "thumbnail_1.jpg", "thumbnail_2.jpg", "thumbnail_3.jpg", "thumbnail_4.jpg"}
	lastName := namesList[4]

	if setNameDigit(namesList) != string(lastName[10]+1) {
		t.Error("Must be string number return!")
	}
}

func TestSetNameDigitZeroList(t *testing.T) {

	namesList := []string{}

	if setNameDigit(namesList) != "0" {
		t.Error("Must be zero string number return!")
	}
}

func TestSetNameDigitWrongNumbers(t *testing.T) {

	namesList := []string{"<<", "<<<<<"}

	out := setNameDigit(namesList)

	if out != "" {
		t.Errorf("Wrong number symbols! %v", out)
	}
}

func TestSetThumbnailNameVaild(t *testing.T) {

	tmb := NewThumbnail()

	result := tmb.setThumbnailName()

	if len(result) == 0 {
		t.Errorf("Empty file name! %s\n", result)
	}
}

func TestNewThumbnailVaild(t *testing.T) {

	videoID := ""
	link := ""
	thumbnailsName := ""

	result := NewThumbnail()

	if result == nil {
		t.Errorf("Should be non nil struct! %v\n", result)
	}

	if result.VideoID != videoID || result.link != link || result.fileName == nil || result.thumbnailsDir == "" || result.thumbnailsName != thumbnailsName {
		t.Errorf("Should be valid struct! %v\n", result)
	}
}

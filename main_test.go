// Run Benchmarks go test -run=XXX -bench=. -benchmem

package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func TestCreateFile(t *testing.T) {

	tmb := &Thumbnail{}

	thumbnailsDir := "./test_data"
	thumbnailsName := "test_data/thumbnail_" + setNameDigit(tmb.fileName) + ".jpg"

	file, err := tmb.createFile(thumbnailsDir, thumbnailsName)
	file.Close()
	defer os.Remove("./" + thumbnailsName)

	if err != nil {
		t.Error("Wrong created file!")
	}
}

func TestCreateFileWrongDirectory(t *testing.T) {

	tmb := &Thumbnail{}
	var thumbnailsDir string

	osName := runtime.GOOS

	if osName == "windows" {
		thumbnailsDir = "/<"
	} else if osName == "linux" {
		thumbnailsDir = ""
	}

	file, _ := tmb.createFile(thumbnailsDir, "thumbnail_")
	file.Close()

	if file != nil {
		t.Error("Expected nil return!")
	}
}

func TestCreateFileWrong(t *testing.T) {

	tmb := &Thumbnail{}

	var thumbnailsDir = "./test_data"
	var thumbnailsName string
	osName := runtime.GOOS

	if osName == "windows" {
		thumbnailsName = "test_data/<"
	} else if osName == "linux" {
		thumbnailsName = "/"
	}

	file, err := tmb.createFile(thumbnailsDir, thumbnailsName)
	file.Close()
	defer os.Remove("./" + thumbnailsName)

	if err == nil {
		t.Error("incorrect file name. Must be nil return!")
	}
}

func TestWriteFile(t *testing.T) {

	file, _ := os.Create("./test_data/thumbnail_test2.jpg")
	defer file.Close()
	defer os.Remove("./test_data/thumbnail_test2.jpg")

	resp, err := http.Get("https://www.youtube.com/watch?v=N2wJQSBx5i4")

	if writeFile(file, resp) != nil {

		t.Errorf("Write file failed %v\n", err)
	}
	resp.Body.Close()
}

func TestWriteFileWrong(t *testing.T) {

	resp, _ := http.Get("https://www.youtube.com/watch?v=N2wJQSBx5i4")
	err := writeFile(nil, resp)

	if err == nil {
		t.Errorf("Write file failed %v\n", err)
	}
	resp.Body.Close()
}

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

func TestGetURLResponse(t *testing.T) {

	idList := []string{"N2wJQSBx5i4", "65AB2pMCj4I", "ZnhquCll3uQ"}

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

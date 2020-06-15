package main

import (
	"net/http"
	"os"
	"runtime"
	"testing"
)

const TestDir = "./test_data"

func TestCreateFile(t *testing.T) {

	tmb := &Thumbnail{}

	// thumbnailsDir := "./test_data"
	_ = createFolder(TestDir)
	thumbnailsName := TestDir[2:] + "/thumbnail_" + setNameDigit(tmb.fileName) + ".jpg"

	file, err := createFile(thumbnailsName)
	if err != nil {
		t.Error("Wrong created file!")
	}
	file.Close()
	os.Remove("./" + thumbnailsName)
}

func TestCreateWrongDirectory(t *testing.T) {

	var thumbnailsDir string

	osName := runtime.GOOS

	if osName == "windows" {
		thumbnailsDir = "/<"
	} else if osName == "linux" {
		thumbnailsDir = ""
	}

	err := createFolder(thumbnailsDir)
	if err == nil {
		t.Error("Expected nil return!")
	}
}

func TestCreateFileWrong(t *testing.T) {

	// thumbnailsDir := "./test_data"
	_ = createFolder(TestDir)
	var thumbnailsName string
	osName := runtime.GOOS

	if osName == "windows" {
		thumbnailsName = TestDir[2:] + "/<"
	} else if osName == "linux" {
		thumbnailsName = "/"
	}

	file, err := createFile(thumbnailsName)
	file.Close()
	defer os.Remove("./" + thumbnailsName)

	if err == nil {
		t.Error("incorrect file name. Must be nil return!")
	}
}

func TestWriteFile(t *testing.T) {

	file, _ := os.Create(TestDir + "/thumbnail_test2.jpg")
	defer file.Close()
	defer os.Remove(TestDir + "/thumbnail_test2.jpg")

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

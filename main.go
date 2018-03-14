// Remove Symbol and Debug info at compile
// go build -ldflags "-s -w"

// Run on linux and mac os x machines
// sudo ./get-thumbnail-Youtube

package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

// Thumbnail struct has a three properties whose get the correct url and current file names.
type Thumbnail struct {
	VideoID  string
	link     string
	fileName []string
}

func main() {

	t := &Thumbnail{}

	fmt.Printf("\nEnter Youtube Url: ")
	fmt.Scanf("%s ", &t.link)

	t.findVideoID(t.link)

	const thumbnailsDir = "./thumbnails"

	err := createFolder(thumbnailsDir)
	checkErr(err)

	// tree walks thumbnails directory
	filepath.Walk(thumbnailsDir, t.walkFunc)

	thumbnailsName := "thumbnails/thumbnail_" + setNameDigit(t.fileName) + ".jpg"
	readyFile, errCreate := t.createFile(thumbnailsName)
	checkErr(errCreate)

	errWrite := writeFile(readyFile, t.getURLResponse())
	checkErr(errWrite)
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

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

func createFolder(thumbnailsDir string) error {

	// create folder if already exist do nothing
	err := os.MkdirAll(thumbnailsDir, os.ModePerm)
	if err != nil {
		return errors.New("Can't create thumbnails folder")
	}
	return nil
}

// createFile create and save jpg thumbnail file.
// Default folder for file is "thumbnails" at the root directory.
func (t *Thumbnail) createFile(thumbnailsName string) (*os.File, error) {

	// create file with auto set in the name's last number
	createdFile, err := os.Create(thumbnailsName)

	if err != nil {
		log.Println(err)
		return nil, errors.New("Can't create file")
	}

	return createdFile, nil
}

// writeFile write response body from valid url at the created jpg thumbnail file.
func writeFile(readyFile *os.File, resp *http.Response) error {

	_, err := io.Copy(readyFile, resp.Body)
	if err != nil {
		log.Println(err)
		return err
	}
	resp.Body.Close()
	readyFile.Close()
	fmt.Println("\nAlready Done:)")
	return nil
}

// findVideoID extract video id from raw input url and save it at Thumbnail struct.
// Also checks id length and bad symbols at id sequence.
func (t *Thumbnail) findVideoID(urlVideo string) error {

	equalIndex := strings.Index(urlVideo, "=")
	ampIndex := strings.Index(urlVideo, "&")
	slash := strings.LastIndex(urlVideo, "/")
	questionIndex := strings.Index(urlVideo, "?")
	var id string

	if equalIndex != -1 {

		if ampIndex != -1 {
			id = urlVideo[equalIndex+1 : ampIndex]
		} else if questionIndex != -1 && strings.Contains(urlVideo, "?t=") {
			id = urlVideo[slash+1 : questionIndex]
		} else {
			id = urlVideo[equalIndex+1:]
		}

	} else {
		id = urlVideo[slash+1:]
	}

	t.VideoID = id

	if strings.ContainsAny(id, "?&/<%=") {
		return errors.New("invalid characters in video id")
	}
	if len(id) < 10 {
		return errors.New("the video id must be at least 10 characters long")
	}
	return nil
}

// walkFunc checks root directory name for persist thumbnails.
// If name does not exist, walkFunc save new name into Thumbnail struct.
func (t *Thumbnail) walkFunc(path string, info os.FileInfo, err error) error {

	if "thumbnails" != info.Name() {
		t.fileName = append(t.fileName, info.Name())
	}
	return nil
}

// setNameDigit gets last file name at thumbnails directory and
// sets next filenames number at string representation.
func setNameDigit(inputArr []string) string {

	if len(inputArr) != 0 {

		// sort and get last thumbnail filename
		sort.Strings(inputArr)
		lastFile := inputArr[len(inputArr)-1]

		// stackoverflow fast and optimal solution
		var numbers string
		for i := 0; i < len(lastFile); i++ {

			elem := lastFile[i]
			if '0' <= elem && elem <= '9' {
				numbers += strconv.Itoa(int(elem) - '0')
			}
		}

		digitsCounter, err := strconv.Atoi(numbers)
		if err != nil {
			fmt.Println("String to Int Atoi conversion error!", err)
			return ""
		}

		digitsCounter++
		return strconv.Itoa(digitsCounter)
	}
	return "0"
}

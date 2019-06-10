package main

import (
	"errors"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Thumbnail struct has a three properties whose get the correct url and current file names.
type Thumbnail struct {
	VideoID        string
	link           string
	fileName       []string
	thumbnailsDir  string
	thumbnailsName string
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

	if len(inputArr) > 0 {

		// sort and get last thumbnail filename
		sort.Slice(inputArr, func(i, j int) bool {

			startA := strings.Index(inputArr[i], "_") + 1
			endA := strings.Index(inputArr[i], ".")

			startB := strings.Index(inputArr[j], "_") + 1
			endB := strings.Index(inputArr[j], ".")

			if startA == -1 || endA == -1 || startB == -1 || endB == -1 {
				return false
			}

			numberA := inputArr[i][startA:endA]
			numberB := inputArr[j][startB:endB]

			numA, _ := strconv.Atoi(numberA)
			numB, _ := strconv.Atoi(numberB)

			return numA < numB
		})

		lastFile := inputArr[len(inputArr)-1]

		var numbers string
		for i := 0; i < len(lastFile); i++ {

			elem := lastFile[i]
			if '0' <= elem && elem <= '9' {
				numbers += strconv.Itoa(int(elem) - '0')
			}
		}

		digitsCounter, err := strconv.Atoi(numbers)
		if err != nil {
			log.Println("String to Int Atoi conversion error!", err)
			return ""
		}

		digitsCounter++
		return strconv.Itoa(digitsCounter)
	}
	return "0"
}

func (t *Thumbnail) setThumbnailName() string {
	return t.thumbnailsDir + "/thumbnail_" + setNameDigit(t.fileName) + ".jpg"
}

func NewThumbnail() *Thumbnail {
	return &Thumbnail{
		VideoID:        "",
		link:           "",
		fileName:       []string{},
		thumbnailsDir:  "./thumbnails",
		thumbnailsName: "",
	}
}

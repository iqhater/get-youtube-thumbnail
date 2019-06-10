// Remove Symbol and Debug info at compile
// go build -ldflags "-s -w"

// Run on linux and mac os x machines
// sudo ./get-thumbnail-Youtube

package main

import (
	"fmt"
	"log"
	"path/filepath"
)

func main() {

	t := NewThumbnail()

	fmt.Printf("\nEnter Youtube Url: ")
	fmt.Scanf("%s ", &t.link)

	err := t.findVideoID(t.link)
	checkErr(err)

	err = createFolder(t.thumbnailsDir)
	checkErr(err)

	// Walk walks thumbnails dir and save file names
	err = filepath.Walk(t.thumbnailsDir, t.walkFunc)
	checkErr(err)

	readyFile, errCreate := createFile(t.setThumbnailName())
	checkErr(errCreate)

	errWrite := writeFile(readyFile, t.getURLResponse())
	checkErr(errWrite)
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

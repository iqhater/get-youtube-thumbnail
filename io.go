package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// createFolder function is create a thumbnails folder
// with a given name as parametr
// Folder will be stored all thumbnails
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
func createFile(thumbnailsName string) (*os.File, error) {

	// create file with auto set in the name's last number
	createdFile, err := os.Create(thumbnailsName)
	if err != nil {
		log.Println(err)
		return nil, errors.New("Can't create file")
	}

	return createdFile, nil
}

// writeFile write response body from valid url
// at the created jpg thumbnail file.
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

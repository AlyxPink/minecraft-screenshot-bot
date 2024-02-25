package screenshot

import (
	"io/fs"
	"os"
	"path/filepath"
	"time"

	"github.com/charmbracelet/log"
	"github.com/google/uuid"
)

type Screenshot struct {
	ID uuid.UUID

	Description string
	File        []byte
	Name        string
	Path        string
	URL         string
}

func GetLatest() *Screenshot {
	fileInfo, err := getLastCreatedFile()
	if err != nil {
		log.Fatal(err)
	}

	if fileInfo != nil {
		log.Info("Screenshot found: ", "name", fileInfo.Name(), "createdAt", fileInfo.ModTime())
	} else {
		log.Fatal("No screenshots found", "SCREENSHOTS_DIR_PATH", os.Getenv("SCREENSHOTS_DIR_PATH"))
	}

	filePath := filepath.Join(os.Getenv("SCREENSHOTS_DIR_PATH"), fileInfo.Name())

	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal("Error while opening screenshot file", "path", filePath, "error", err)
	}

	return &Screenshot{
		ID:   uuid.New(),
		Name: fileInfo.Name(),
		Path: filePath,
		File: file,
	}
}

// getLastCreatedFile takes a directory path as an argument and returns the FileInfo
// of the most recently created file in that directory. If the directory is empty,
// or if there are no files in the directory, it returns nil.
func getLastCreatedFile() (fs.FileInfo, error) {
	dir := os.Getenv("SCREENSHOTS_DIR_PATH")
	// Initialize variables to store information about the newest file
	var newestFile fs.FileInfo
	var newestTime time.Time

	// Walk through the directory and its subdirectories
	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		// If an error occurs, return the error and stop walking the directory
		if err != nil {
			return err
		}
		// If the current entry is a directory, skip it
		if d.IsDir() {
			return nil
		}

		// Retrieve the FileInfo of the current file
		info, err := d.Info()
		if err != nil {
			return err
		}

		// Check if the current file is newer than the newest file found so far
		if info.ModTime().After(newestTime) {
			// If it is, update newestFile and newestTime with the current file's information
			newestFile = info
			newestTime = info.ModTime()
		}
		// Continue walking through the directory
		return nil
	})

	// If an error occurred during the directory walk, return the error
	if err != nil {
		return nil, err
	}

	// Return the FileInfo of the newest file found (or nil if no files were found)
	return newestFile, nil
}

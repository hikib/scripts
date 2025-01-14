package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {
	// get DOWNLOADS directory
	downDirectory, err := getDownloadDirectory()
	if err != nil {
		fmt.Println("Error getting most 'Downloads' directory:", err)
		os.Exit(1)
	}

	// get most recently modified FILE in DOWNLOADS directory
	downFile, err := getMostRecentFile(downDirectory)
	if err != nil {
		fmt.Println("Error getting most recent downloaded file:", err)
		os.Exit(1)
	}

	var destination string
	if len(os.Args) > 1 {
		// get SELECTED directory
        dirPath := os.Args[1]
        info, err := os.Stat(dirPath)
		if err != nil {
			fmt.Println("Error with given path:", err)
			os.Exit(1)
		}
        if !info.IsDir() {
			fmt.Println("Given path is not a valid directory")
			os.Exit(1)
        }
		destination = filepath.Join(os.Args[1], filepath.Base(downFile))

	} else {
		// get CURRENT directory
		currentDir, err := os.Getwd()
		if err != nil {
			fmt.Println("Error getting current directory:", err)
			os.Exit(1)
		}
		destination = filepath.Join(currentDir, filepath.Base(downFile))
	}

	// move FILE to DESTINATION
	err = os.Rename(downFile, destination)
	if err != nil {
		fmt.Println("Error moving file:", err)
		os.Exit(1)
	}
}

func getDownloadDirectory() (downloadsPath string, err error) {
	downloadsPath = os.Getenv("DOWNLOADS")
	if downloadsPath == "" {
		return "", fmt.Errorf("DOWNLOADS environment variable not set or empty")
	}

	return downloadsPath, nil
}

func getMostRecentFile(directory string) (mostRecentFile string, err error) {
	files, err := os.ReadDir(directory)
	if err != nil {
		return "", fmt.Errorf("Could not read directory: %v", err)
	}
	if len(files) == 0 {
		return "", fmt.Errorf("No files found in directory.")
	}

	// limit search to files within last 5 min.
	// "NEW" files
	threshold := 5 * time.Minute
	currentTime := time.Now()
	var mostRecentTime time.Time

	for _, file := range files {
		// Skip directories
		if file.IsDir() {
			continue
		}
		info, err := file.Info()
		if err != nil {
			continue
		}
		modTime := info.ModTime()
		if currentTime.Sub(modTime) > threshold {
			continue
		}

		if modTime.After(mostRecentTime) {
			mostRecentFile = filepath.Join(directory, file.Name())
			mostRecentTime = modTime
		}
	}

	if mostRecentFile == "" {
		return "", fmt.Errorf("No new file in directory.")
	}
	return mostRecentFile, nil
}

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// Handles my files while writing job applications.
// Moves the temporary files from 'docs' to the archive 'apps'.
func main() {
	repoPath, companyName := parseInput()
	archive := filepath.Join(repoPath, "apps") // TODO: Exists?
	docs := filepath.Join(repoPath, "docs")    // TODO: Exists?

	dirName := getNewDirName(companyName)
	dirPath := filepath.Join(archive, dirName)
	os.Mkdir(dirPath, 0777) // TODO: Handle ERROR

	files := map[string]string{
		"ansøgning.pdf": "coverletter/output/coverletter.pdf",
		"cv.pdf":        "cv/output/cv.pdf",
	}

	for name, file := range files {
		fromPath := filepath.Join(docs, file)
		toPath := filepath.Join(dirPath, name)
		// TODO: Handle ERROR
		os.Rename(fromPath, toPath)
	}
}

// Get path to 'job search directory' and company name from arguments
func parseInput() (string, string) {
	repoPath, _ := filepath.Abs(os.Args[1]) // TODO: Error?
	companyName := os.Args[2]
	return repoPath, companyName
}

// Combining current year, week number and company name
func getNewDirName(companyName string) string {
	yearNum, weekNum := time.Now().ISOWeek()
	week := fmt.Sprintf("%02d", weekNum)
	year := strconv.Itoa(yearNum)
	return strings.Join([]string{year, week, companyName}, "-")
}

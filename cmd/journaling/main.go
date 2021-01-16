package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const (
	subtitle   string = "## Reviews"
	reviewPath string = "../pipelines/reviews/"
	fileExt    string = ".md"
)

// Prints the links to review files in markdown format to the console.
// The output is used in vimwiki as a template for the journal.
func main() {
	filePath := os.Args[1]
	date := FileToDate(filepath.Base(filePath))
	entry := JournalEntry{date}
	fmt.Print(entry)
}

// Creates a time.Time object
// from a file name in vimwiki diary format '2020-01-01.md'
func FileToDate(fileName string) time.Time {
	name := strings.TrimSuffix(fileName, fileExt)
	dateParts := strings.Split(name, "-")

	year, _ := strconv.Atoi(dateParts[0])
	month, _ := strconv.Atoi(dateParts[1])
	day, _ := strconv.Atoi(dateParts[2])

	date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	return date
}

type JournalEntry struct {
	date time.Time
}

func (e JournalEntry) String() string {
	year, month, day := e.date.Date()
	_, weekNum := e.date.ISOWeek()
	weekday := e.date.Weekday()

	titel := fmt.Sprintf(
		"# %02d.%s %d, Week %02d, %s\n",
		day,
		month,
		year,
		weekNum,
		weekday)

	weekLink := fmt.Sprintf(
		"- [Week %d](%s%d-W%02d%s)",
		weekNum,
		reviewPath,
		year,
		weekNum,
		fileExt)

	monthLink := fmt.Sprintf(
		"- [%s](%s%d-M%02d%s)",
		month,
		reviewPath,
		year,
		month,
		fileExt)

	yearLink := fmt.Sprintf(
		"- [%d](%s%d%s)\n",
		year,
		reviewPath,
		year,
		fileExt)

	return strings.Join(
		[]string{
			titel,
			subtitle,
			weekLink,
			monthLink,
			yearLink}, "\n")
}

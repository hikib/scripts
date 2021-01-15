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

	entry := Entry{date}
	fmt.Print(entry)
}

func FileToDate(fileName string) time.Time {
	name := strings.TrimSuffix(fileName, fileExt)
	split := strings.Split(name, "-")

	year, _ := strconv.Atoi(split[0])
	month, _ := strconv.Atoi(split[1])
	day, _ := strconv.Atoi(split[2])

	date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	return date
}

type Entry struct {
	date time.Time
}

func (e Entry) String() string {
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
		"- [Week %d](%s%d-W%02d.md)",
		weekNum,
		reviewPath,
		year,
		weekNum)

	monthLink := fmt.Sprintf(
		"- [%s](%s%d-M%02d.md)",
		month,
		reviewPath,
		year,
		month)

	yearLink := fmt.Sprintf(
		"- [%d](%s%d.md)\n",
		year,
		reviewPath,
		year)

	return strings.Join([]string{
		titel,
		subtitle,
		weekLink,
		monthLink,
		yearLink}, "\n")
}

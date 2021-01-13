package main

import (
	"fmt"
	"strings"
	"time"
)

const (
	subtitle   string = "## Reviews"
	reviewPath string = "../pipelines/reviews/"
)

// Prints the links to review files in markdown format to the console.
// The output is used in vimwiki as a template for the journal.
func main() {
	entry := Entry{time.Now()}
	fmt.Print(entry)
}

type Entry struct {
	date time.Time
}

func (e Entry) String() string {
	year, month, day := e.date.Date()
	_, weekNum := e.date.ISOWeek()
	weekday := e.date.Weekday()

	titel := fmt.Sprintf("# %s %02d.%s, Week %02d\n", weekday, day, month, weekNum)
	weekLink := fmt.Sprintf("- [Week %d](%s%d-W%02d.md)", weekNum, reviewPath, year, weekNum)
	monthLink := fmt.Sprintf("- [%s](%s%d-M%02d.md)", month, reviewPath, year, month)
	yearLink := fmt.Sprintf("- [%d](%s%d.md)\n", year, reviewPath, year)

	return strings.Join([]string{titel, subtitle, weekLink, monthLink, yearLink}, "\n")
}

package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	subtitle   string = "## Reviews \n"
	reviewPath string = "../pipelines/reviews/"
)

// Prints the date and links to review files in markdown format to the
// console. The output is used in vimwiki as a template for the journal.
func main() {
	yearNum, monthNum, dayNum := time.Now().Date()
	_, weekNum := time.Now().ISOWeek()
	var (
		day       string = strconv.Itoa(dayNum) + "."
		year      string = strconv.Itoa(yearNum)
		week      string = fmt.Sprintf("%02d", weekNum)
		monthName string = fmt.Sprintf("%s", monthNum)
		month     string = fmt.Sprintf("%02d", monthNum)
	)

	date := strings.Join([]string{day, monthName, year}, " ")
	title := fmt.Sprintf("# %s \n\n", date)

	weekFile := strings.Join([]string{year, "W" + week}, "-")
	weekLink := fmt.Sprintf("- [Week](%s%s.md) \n", reviewPath, weekFile)

	monthFile := strings.Join([]string{year, "M" + month}, "-")
	monthLink := fmt.Sprintf("- [Month](%s%s.md) \n", reviewPath, monthFile)

	yearLink := fmt.Sprintf("- [Year](%s%s.md) \n", reviewPath, year)

	template := title +
		subtitle +
		weekLink +
		monthLink +
		yearLink

	fmt.Print(template)
}

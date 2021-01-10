package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	subtitle   string = "## Reviews"
	reviewPath string = "../pipelines/reviews/"
)

// Prints the date and links to review files in markdown format to the
// console. The output is used in vimwiki as a template for the journal.
func main() {
	yearNum, monthNum, dayNum := time.Now().Date()
	_, weekNum := time.Now().ISOWeek()

	// time.Now() objects to strings
	day := strconv.Itoa(dayNum)
	year := strconv.Itoa(yearNum)
	week := fmt.Sprintf("%02d", weekNum)
	monthName := fmt.Sprintf("%s", monthNum)
	month := fmt.Sprintf("%02d", monthNum)

	// Markdown formatted parts of template
	date := getDate(day, monthName, year)
	weekLink := getWeekLink(week, year)
	monthLink := getMonthLink(month, year)
	yearLink := getYearLink(year)

	template := getTemplate(date, weekLink, monthLink, yearLink)

	fmt.Print(template)
}

// Current date
func getDate(day, month, year string) string {
	dateParts := []string{day + "." + month + ",", year}
	return strings.Join(dateParts, " ")
}

// Markdown formatted link to current weeks review file
func getWeekLink(week, year string) string {
	fileName := strings.Join([]string{year, "W" + week}, "-")
	return fmt.Sprintf("- [Week](%s%s.md)", reviewPath, fileName)
}

// Markdown formatted link to current months review file
func getMonthLink(month, year string) string {
	fileName := strings.Join([]string{year, "M" + month}, "-")
	return fmt.Sprintf("- [Month](%s%s.md)", reviewPath, fileName)
}

// Markdown formatted link to current years review file
func getYearLink(year string) string {
	return fmt.Sprintf("- [Year](%s%s.md)", reviewPath, year)
}

// Join markdown formatted parts to whole template file
func getTemplate(date, weekLink, monthLink, yearLink string) string {
	date += "\n"     // Empty newline
	yearLink += "\n" // Empty newline at end
	templateParts := []string{date, subtitle, weekLink, monthLink, yearLink}
	return strings.Join(templateParts, "\n")
}

package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Prints the links to review files in markdown format to the console.
// The output is used in vimwiki as a template for the journal.
func main() {
	yearNum, monthNum, _ := time.Now().Date()
	_, weekNum := time.Now().ISOWeek()

	year := Year{yearNum}
	month := Month{year, monthNum}
	week := Week{year, weekNum}

	// Print template - line by line
	fmt.Println(subtitle)
	fmt.Println(week.MarkdownLink())
	fmt.Println(month.MarkdownLink())
	fmt.Println(year.MarkdownLink())
}

const (
	subtitle   string = "## Reviews"
	reviewPath string = "../pipelines/reviews/"
)

type Link interface {
	MarkdownLink() string
}

type Year struct {
	yearNum int
}

func (y Year) String() string {
	return strconv.Itoa(y.yearNum)
}

func (y Year) MarkdownLink() string {
	return fmt.Sprintf("- [Year](%s%s.md)", reviewPath, y.String())
}

type Month struct {
	year  Year
	month time.Month
}

func (m Month) Num() string {
	return fmt.Sprintf("%02d", m.month)
}

func (m Month) MarkdownLink() string {
	fileName := strings.Join([]string{m.year.String(), "M" + m.Num()}, "-")
	return fmt.Sprintf("- [Month](%s%s.md)", reviewPath, fileName)
}

type Week struct {
	year    Year
	weekNum int
}

func (w Week) String() string {
	return fmt.Sprintf("%02d", w.weekNum)
}

func (w Week) MarkdownLink() string {
	fileName := strings.Join([]string{w.year.String(), "W" + w.String()}, "-")
	return fmt.Sprintf("- [Week](%s%s.md)", reviewPath, fileName)
}

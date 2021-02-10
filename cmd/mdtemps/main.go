package main

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const (
	fileExt string = ".md"
)

type Review struct {
	Path        string
	Name        string
	Period      string
	Date        time.Time
	Description string
}

func (r *Review) Init(path string) {
	r.Path = path
	r.Name = strings.TrimSuffix(filepath.Base(path), fileExt)
	r.Period = r.getPeriod()
	r.Date = r.getDate()
	r.Description = r.composeDescription()
}

func (r *Review) getDate() time.Time {
	dateParts := strings.Split(r.Name, "-")
	year, _ := strconv.Atoi(dateParts[0])

	switch r.Period {
	case "Monthly":
		month, _ := strconv.Atoi(dateParts[1])
		date := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	case "Weekly":
		dateParts := strings.Split(r.Name, "-")
	default:
		return "Resolutions 2021"
	}

	return date
}

func (r *Review) composeDescription() string {
	switch r.Period {
	case "Monthly":
		return "January"
	case "Weekly":
		return "Week 01"
	default:
		return "Resolutions 2021"
	}
}

func (r *Review) getPeriod() string {
	switch {
	case strings.Contains(r.Name, "M"):
		return "Monthly"
	case strings.Contains(r.Name, "W"):
		return "Weekly"
	default:
		return "Annual"
	}
}

func main() {
	// Reviews
	// 2021-W04.md
	// 2021-M02.md
	// 2021.md
	filePath := "~/vimwiki/pipelines/reviews/2021-M01.md"

	review := Review{}
	review.Init(filePath)
	fmt.Println(review.Path)
	fmt.Println(review.Name)
	fmt.Println(review.Period)
}

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	title     string = "# Reviews\n"
	indexFile string = "README.md"
	fileExt   string = ".md"
)

// TODO: Group files in years, then months. Similar to DiaryIndex file
func main() {
	fmt.Println(title)

	files := GetFiles(os.Args[1])
	for _, f := range files {
		if isReviewFile(f) {
			fmt.Println(GetMarkdownLink(f.Name()))
		}
	}
}

func GetMarkdownLink(fileName string) string {
	name := strings.TrimSuffix(fileName, fileExt)
	var descr string

	switch {
	case strings.Contains(name, "W"):
		num := strings.Split(name, "-W")[1]
		descr = fmt.Sprintf("WEEK %s", num)
	case strings.Contains(name, "M"):
		num := strings.Split(name, "-M")[1]
		month, _ := strconv.Atoi(num)
		descr = time.Month(month).String()
	default:
		descr = "New Years Resolution"
	}

	return fmt.Sprintf("- [%s](%s)", descr, fileName)
}

func GetFiles(filePath string) []os.FileInfo {
	files, err := ioutil.ReadDir(filepath.Dir(filePath))
	if err != nil {
		log.Fatal(err)
	}

	sort.SliceStable(files, func(i, j int) bool {
		return files[i].Name() > files[j].Name()
	})

	return files
}

func isReviewFile(f os.FileInfo) bool {
	var validFile = regexp.MustCompile(`^20\d{2}(-[MW][0-5]\d)*\.md$`)
	if !f.IsDir() && validFile.MatchString(f.Name()) {
		return true
	}
	return false
}

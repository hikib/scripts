package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

const (
	title     string = "# Reviews\n"
	indexFile string = "README.md"
	fileExt   string = ".md"
)

func main() {
	// entries := map[string][]os.FileInfo{}
	filePath := os.Args[1]
	files := GetFiles(filePath)

	// fmt.Print(fmt.Sprintf("%v", entries))

	fmt.Println(title)
	for _, f := range files {
		fileName := f.Name()
		if !(fileName == indexFile || f.IsDir()) && strings.HasSuffix(fileName, fileExt) {
			// fmt.Println(fileName)
			// name := strings.TrimSuffix(fileName, fileExt)
			fmt.Println(GetMarkdownLink(fileName))
		}
	}
}

type Entry struct {
	year  string
	files []os.FileInfo
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
		descr = fmt.Sprintf("MONTH %s", num)
	default:
		descr = "New Years Resolution"
	}

	return fmt.Sprintf("[%s](%s)", descr, fileName)
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

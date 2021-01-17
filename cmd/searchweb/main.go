package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"strings"
)

func main() {
	pagePtr := flag.String(
		"page",
		"duckduckgo",
		"ordnet, korpus, tysk, github, youtube")
	flag.Parse()

	var url string
	switch *pagePtr {
	case "duck":
		url = "http://duckduckgo.com/?q="
	case "ordnet":
		url = "https://ordnet.dk/ddo/ordbog?query="
	case "korpus":
		url = "https://ordnet.dk/korpusdk/quick_search?SearchableText="
	case "tysk":
		url = "https://de.langenscheidt.com/daenisch-deutsch/"
	case "youtube":
		url = "https://www.youtube.com/results?search_query="
	case "github":
		url = "https://github.com/search?q="
	}

	query := strings.Join(flag.Args(), "+")
	openBrowser(url + query)
}

func openBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"flag"
	"log"
	"strings"

	"github.com/hikmet-kibar/scripts/cmd/searchweb/browser"
	"github.com/hikmet-kibar/scripts/cmd/searchweb/config"
)

var pagePtr = flag.String(
	"page",
	"duck",
	"Pages set in config.yaml.")

var configPtr = flag.String(
	"config",
	"",
	"Absolute path to your *.yaml file with pages OR environment variable 'PAGES'")

func main() {
	flag.Parse()

	cfg, err := config.GetConfig(*configPtr)
	if err != nil {
		log.Fatal(err)
	}

	url := cfg.Pages[*pagePtr]
	query := strings.Join(flag.Args(), "+")
	browser.Open(url + query)
}

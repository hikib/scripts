package main

import (
	"flag"
	"log"
	"strings"

	"github.com/hikmet-kibar/searchweb/browser"
	"github.com/hikmet-kibar/searchweb/config"
)

var pagePtr = flag.String(
	"page",
	"duck",
	"Pages set in config.yaml.")

var configPtr = flag.String(
	"config",
	"config.yaml",
	"Absolute path to your config.yaml file.")

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

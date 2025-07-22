package main

import (
	"fmt"
	"log"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {
	defer cfg.wg.Done()

	parsedURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - crawlPage: couldn't parse URL '%s': %v\n", rawCurrentURL, err)
		return
	}

	if parsedURL.Hostname() != cfg.baseURL.Hostname() {
		return
	}

	currentURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		log.Printf("error normalising url: %v", err)
		return
	}

	if !cfg.addPageVisit(currentURL) {
		log.Printf("already visited url: %s", currentURL)
		return
	}

	fmt.Println(currentURL)

	data, err := getHTML(currentURL)
	if err != nil {
		fmt.Printf("error getting getting html from url: %v", err)
		return
	}

	urls, err := getURLsFromHTML(data, cfg.baseURL)
	if err != nil {
		log.Printf("error reading urls from html: %v", err)
		return
	}

	for _, url := range urls {
		cfg.wg.Add(1)
		go cfg.crawlPage(url)
	}
}

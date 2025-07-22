package main

import (
	"fmt"
	"log"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {

	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()

	if cfg.pagesLens() >= cfg.maxPages {
		return
	}

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
		// log.Printf("already visited url: %s", currentURL)
		return
	}

	// fmt.Println(currentURL)

	htmlData, err := getHTML(currentURL)
	if err != nil {
		log.Printf("error getting getting html from url: %v", err)
		return
	}

	urls, err := getURLsFromHTML(htmlData, cfg.baseURL)
	if err != nil {
		log.Printf("error reading urls from html: %v", err)
		return
	}

	for _, nextURL := range urls {
		cfg.wg.Add(1)
		// log.Printf("Spawning crawl for %s", nextURL)
		go cfg.crawlPage(nextURL)
	}
}

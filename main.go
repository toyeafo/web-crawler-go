package main

import (
	"errors"
	"fmt"
	"log"
	"net/url"
	"os"
	"sync"
)

type config struct {
	pages              map[string]int
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
}

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println(errors.New("no website provided"))
		os.Exit(1)
	}
	if len(args) > 2 {
		fmt.Println(errors.New("too many arguments provided"))
		os.Exit(1)
	}

	rawBaseURL := args[1]

	u, err := url.Parse(rawBaseURL)
	if err != nil {
		log.Fatalf("error parsing url: %v, error: %v", rawBaseURL, err)
	}

	maxConcurrency := 3

	cfg := config{
		pages:              make(map[string]int),
		baseURL:            u,
		concurrencyControl: make(chan struct{}, maxConcurrency),
		mu:                 &sync.Mutex{},
		wg:                 &sync.WaitGroup{},
	}

	cfg.wg.Add(1)
	go cfg.crawlPage(cfg.baseURL.String())
	cfg.wg.Wait()

	for k, v := range cfg.pages {
		fmt.Printf("URL: %s, Visited: %d\n", k, v)
	}
}

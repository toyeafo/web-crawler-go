package main

import (
	"errors"
	"fmt"
	"os"
)

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

	const maxConcurrency = 3
	cfg, err := configure(rawBaseURL, maxConcurrency)
	if err != nil {
		fmt.Printf("Error configuring: %v", err)
		return
	}

	cfg.wg.Add(1)
	go cfg.crawlPage(rawBaseURL)
	cfg.wg.Wait()

	for k, v := range cfg.pages {
		fmt.Printf("URL: %s, Visited: %d\n", k, v)
	}
}

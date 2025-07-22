package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	if len(args) < 4 {
		fmt.Println(errors.New("no website provided"))
		os.Exit(1)
	}
	if len(args) > 4 {
		fmt.Println(errors.New("too many arguments provided"))
		os.Exit(1)
	}

	rawBaseURL := args[1]
	maxConcurrencyString := os.Args[2]
	maxPagesString := os.Args[3]

	maxConcurrency, err := strconv.Atoi(maxConcurrencyString)
	if err != nil {
		fmt.Printf("Error - maxConcurrency: %v", err)
		return
	}
	maxPages, err := strconv.Atoi(maxPagesString)
	if err != nil {
		fmt.Printf("Error - maxPages: %v", err)
		return
	}

	cfg, err := configure(rawBaseURL, maxConcurrency, maxPages)
	if err != nil {
		fmt.Printf("Error configuring: %v", err)
		return
	}

	cfg.wg.Add(1)
	go cfg.crawlPage(rawBaseURL)
	cfg.wg.Wait()

	printReport(cfg.pages, rawBaseURL)
}

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

	req, err := getHTML(rawBaseURL)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s", req)
}

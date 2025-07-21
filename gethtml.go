package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	res, err := http.Get(rawURL)
	if res.StatusCode > 399 {
		return "", fmt.Errorf("invalid request")
	}

	defer res.Body.Close()

	if !strings.Contains(res.Header.Get("content-type"), "text/html") {
		return "", fmt.Errorf("invalid html")
	}
	if err != nil {
		return "", fmt.Errorf("error performing request: %w", err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %w", err)
	}
	return string(body), nil
}

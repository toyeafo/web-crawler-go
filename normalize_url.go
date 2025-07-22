package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(rawURL string) (string, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return "", fmt.Errorf("error parsing url: %v, error: %w", rawURL, err)
	}
	parsed_url := strings.ToLower(fmt.Sprintf("%v://%v", u.Scheme, u.Host+u.Path))
	parsed_url = strings.TrimSuffix(parsed_url, "/")
	return parsed_url, nil
}

package main

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	base, err := url.Parse(rawBaseURL)
	if err != nil {
		return []string{}, fmt.Errorf("invalid base URL: %w", err)
	}

	doc, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		return []string{}, fmt.Errorf("error parsing html body: %w", err)
	}

	var links []string

	for n := range doc.Descendants() {
		if n.Type == html.ElementNode && n.DataAtom == atom.A {
			for _, a := range n.Attr {
				if a.Key == "href" {
					u, err := url.Parse(a.Val)
					if err != nil {
						fmt.Printf("error parsing url: %v", err)
						continue
					}

					urlresolved := base.ResolveReference(u).String()
					links = append(links, urlresolved)
					break
				}
			}
		}
	}

	return links, nil
}

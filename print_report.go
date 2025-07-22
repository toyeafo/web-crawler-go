package main

import (
	"fmt"
	"sort"
)

type PageVisit struct {
	Url    string
	Visits int
}

func printReport(pages map[string]int, baseURL string) {
	fmt.Println("===============================================")
	fmt.Printf("           REPORT for %s\n", baseURL)
	fmt.Println("===============================================")

	sortedPages := sortPages(pages)

	for _, val := range sortedPages {
		fmt.Printf("Found %d internal links to %s\n", val.Visits, val.Url)
	}
}

func sortPages(pages map[string]int) []PageVisit {
	var pageVisits []PageVisit

	for k, v := range pages {
		pageVisits = append(pageVisits, PageVisit{
			Url:    k,
			Visits: v,
		})
	}

	sort.Slice(pageVisits, func(i, j int) bool {
		if pageVisits[i].Visits == pageVisits[j].Visits {
			return pageVisits[i].Url < pageVisits[j].Url
		}
		return pageVisits[i].Visits > pageVisits[j].Visits
	})
	return pageVisits
}

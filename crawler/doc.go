/*
Crawler need you pass a Resty client and return a goquery document.

Dependency:

	go get github.com/PuerkitoBio/goquery
	go get github.com/go-resty/resty/v2

Example:

	package main

	import (
		"fmt"
		"net/url"

		"github.com/fatindeed/go-utils/crawler"
		"github.com/go-resty/resty/v2"
	)

	func main() {
		client := crawler.NewGoqueryClient(resty.New())
		doc1, err := s.GetDoc("https://go.dev/")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(doc1)
		query := make(url.Values)
		query.Add("key", "value")
		doc2, err := s.GetDocWithQuery("https://go.dev/", query)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(doc2)
	}
*/
package crawler

package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func getListing(url string) []string {
	var links []string

	// Create the HTTP client.
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}

	// Set the headers.
	request.Header.Set("pragma", "no-cache")
	request.Header.Set("cache-control", "no-cache")
	request.Header.Set("dnt", "1")
	request.Header.Set("upgrade-insecure-requests", "1")
	request.Header.Set("referer", "https://www.takealot.com")

	resp, _ := client.Do(request)
	if resp.StatusCode == 200 {
		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			fmt.Println(err)
		}

		doc.Find(".product-anchor h4").Each(func(i int, s *goquery.Selection) {
			link := s.Text() //.Attr("href")
			fmt.Println(link)
			link = url + link

			if strings.Contains(link, "biz/") {
				text := s.Text()
				if text != "" && text != "more" {
					links = append(links, link)
				}
			}
		})
	} else {
		fmt.Println("couldn't get page." + resp.Status)
	}

	return links
}

func main() {
	m := getListing("https://www.takealot.com/")
	fmt.Println(strings.Join(m, "\n"))
}

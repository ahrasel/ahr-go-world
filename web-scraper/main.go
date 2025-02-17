package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

func main() {
	// set url for scraping
	url := "https://www.bbc.com/news/articles/clyzy300vlzo"

	// get the html content of the url
	htmlContent := getHTMLContent(url)

	// get the title of the page
	title := getTitle(htmlContent)

	// get the links of the page
	links := getImageLinks(htmlContent)

	// print the title of the page
	fmt.Println("Title: ", title)
	fmt.Println("Links: ")
	// print the links of the page
	for _, link := range links {
		fmt.Println(link)
	}
}

func getHTMLContent(url string) string {
	// get the html content of the url
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to get URL: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	return string(body)
}

func getTitle(htmlContent string) string {
	// Use regex to find the title within the specified XPath
	re := regexp.MustCompile(`<h1[^>]*>(.*?)</h1>`)
	matches := re.FindStringSubmatch(htmlContent)
	if len(matches) > 1 {
		// Strip HTML tags from the title
		reTag := regexp.MustCompile(`<[^>]*>`)
		title := reTag.ReplaceAllString(matches[1], "")
		return title
	}
	return "No title found"
}

func getImageLinks(htmlContent string) []string {
	// Use regex to find the image links
	re := regexp.MustCompile(`<img[^>]+src="([^">]+)"`)
	matches := re.FindAllStringSubmatch(htmlContent, -1)
	var imageLinks []string
	for _, match := range matches {
		imageLinks = append(imageLinks, match[1])
	}
	// Remove duplicate image links
	uniqueLinks := make(map[string]bool)
	for _, link := range imageLinks {
		uniqueLinks[link] = true
	}

	var result []string
	for link := range uniqueLinks {
		result = append(result, link)
	}

	return result
}

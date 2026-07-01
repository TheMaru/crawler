package main

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getURLsFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		return nil, fmt.Errorf("couldn't parse HTML: %w", err)
	}

	var urls []string

	doc.Find("a[href]").Each(func(_ int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		parsedHref, err := url.Parse(href)
		if err != nil {
			return
		}
		urls = append(urls, baseURL.ResolveReference(parsedHref).String())
	})

	return urls, nil
}

func getImagesFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		return nil, fmt.Errorf("couldn't parse HTML: %w", err)
	}

	var images []string

	doc.Find("img[src]").Each(func(_ int, s *goquery.Selection) {
		src, _ := s.Attr("src")
		parsedSrc, err := url.Parse(src)
		if err != nil {
			return
		}
		images = append(images, baseURL.ResolveReference(parsedSrc).String())
	})

	return images, nil
}

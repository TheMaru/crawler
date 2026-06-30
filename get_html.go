package main

import (
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getHeadingFromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Printf("couldn't parse HTML: %v", err)
		return ""
	}

	if h1 := doc.Find("h1"); h1.Length() > 0 {
		return strings.TrimSpace(h1.First().Text())
	}
	if h2 := doc.Find("h2"); h2.Length() > 0 {
		return strings.TrimSpace(h2.First().Text())
	}

	return ""
}

func getFirstParagraphFromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Printf("couldn't parse HTML: %v", err)
		return ""
	}

	if main := doc.Find("main"); main.Length() > 0 {
		if p := main.Find("p"); p.Length() > 0 {
			return strings.TrimSpace(p.First().Text())
		}
		return ""
	}

	if p := doc.Find("p"); p.Length() > 0 {
		return strings.TrimSpace(p.First().Text())
	}

	return ""
}

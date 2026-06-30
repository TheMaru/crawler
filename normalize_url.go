package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(rawURL string) (string, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", fmt.Errorf("couldn't parse URL: %w", err)
	}

	url := parsedURL.Host + parsedURL.Path

	url = strings.ToLower(url)

	url = strings.TrimSuffix(url, "/")

	fmt.Println(url)
	return url, nil
}

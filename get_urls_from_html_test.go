package main

import (
	"net/url"
	"reflect"
	"testing"
)

func TestGetURLsFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
	}{
		{
			name:      "absolute URL",
			inputURL:  "https://crawler-test.com",
			inputBody: `<html><body><a href="https://crawler-test.com"><span>Boot.dev</span></a></body></html>`,
			expected:  []string{"https://crawler-test.com"},
		},
		{
			name:      "relative URL converted to absolute",
			inputURL:  "https://blog.boot.dev",
			inputBody: `<html><body><a href="/path/one"><span>Boot.dev</span></a></body></html>`,
			expected:  []string{"https://blog.boot.dev/path/one"},
		},
		{
			name:     "multiple links, mixed absolute and relative",
			inputURL: "https://blog.boot.dev",
			inputBody: `<html><body>
				<a href="/path/one">One</a>
				<a href="https://other.com/path/two">Two</a>
			</body></html>`,
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/two"},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			baseURL, err := url.Parse(tc.inputURL)
			if err != nil {
				t.Fatalf("Test %v - '%s' FAIL: couldn't parse input URL: %v", i, tc.name, err)
			}

			actual, err := getURLsFromHTML(tc.inputBody, baseURL)
			if err != nil {
				t.Fatalf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
			}

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - %s FAIL: expected: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}

func TestGetImagesFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
	}{
		{
			name:      "relative image converted to absolute",
			inputURL:  "https://crawler-test.com",
			inputBody: `<html><body><img src="/logo.png" alt="Logo"></body></html>`,
			expected:  []string{"https://crawler-test.com/logo.png"},
		},
		{
			name:      "absolute image URL",
			inputURL:  "https://crawler-test.com",
			inputBody: `<html><body><img src="https://cdn.other.com/pic.jpg"></body></html>`,
			expected:  []string{"https://cdn.other.com/pic.jpg"},
		},
		{
			name:     "multiple images, one missing src attribute",
			inputURL: "https://crawler-test.com",
			inputBody: `<html><body>
				<img alt="no source here">
				<img src="/a.png">
				<img src="/b.png">
			</body></html>`,
			expected: []string{"https://crawler-test.com/a.png", "https://crawler-test.com/b.png"},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			baseURL, err := url.Parse(tc.inputURL)
			if err != nil {
				t.Fatalf("Test %v - '%s' FAIL: couldn't parse input URL: %v", i, tc.name, err)
			}

			actual, err := getImagesFromHTML(tc.inputBody, baseURL)
			if err != nil {
				t.Fatalf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
			}

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - %s FAIL: expected: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}

package main

import "testing"

func TestGetHeadingFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputBody string
		expected  string
	}{
		{
			name:      "basic h1",
			inputBody: "<html><body><h1>Test Title</h1></body></html>",
			expected:  "Test Title",
		},
		{
			name:      "fallback to h2 when no h1",
			inputBody: "<html><body><h2>Subheading</h2></body></html>",
			expected:  "Subheading",
		},
		{
			name:      "h1 wins over h2",
			inputBody: "<html><body><h2>Subheading</h2><h1>Main Title</h1></body></html>",
			expected:  "Main Title",
		},
		{
			name:      "no heading returns empty string",
			inputBody: "<html><body><p>Just a paragraph.</p></body></html>",
			expected:  "",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := getHeadingFromHTML(tc.inputBody)
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected: %q, actual: %q", i, tc.name, tc.expected, actual)
			}
		})
	}
}

func TestGetFirstParagraphFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputBody string
		expected  string
	}{
		{
			name: "main paragraph priority",
			inputBody: `<html><body>
				<p>Outside paragraph.</p>
				<main>
					<p>Main paragraph.</p>
				</main>
			</body></html>`,
			expected: "Main paragraph.",
		},
		{
			name:      "first p when no main",
			inputBody: "<html><body><p>First.</p><p>Second.</p></body></html>",
			expected:  "First.",
		},
		{
			name: "first p inside main when several",
			inputBody: `<html><body><main>
				<p>First in main.</p>
				<p>Second in main.</p>
			</main></body></html>`,
			expected: "First in main.",
		},
		{
			name:      "no paragraph returns empty string",
			inputBody: "<html><body><h1>Only a heading</h1></body></html>",
			expected:  "",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := getFirstParagraphFromHTML(tc.inputBody)
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected: %q, actual: %q", i, tc.name, tc.expected, actual)
			}
		})
	}
}

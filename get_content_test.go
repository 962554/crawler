package main

import (
	"testing"
)

var tests = []struct {
	name     string
	input    string
	htmlTag  string
	expected string
}{
	{
		name:     "h1",
		input:    "<html><body><h1>Test Title</h1></body></html>",
		htmlTag:  "h1",
		expected: "Test Title",
	},
	{
		name: "p from main",
		input: `<html><body>
		<p>Outside paragraph.</p>
		<main>
			<p>Main paragraph.</p>
		</main>
	</body></html>`,
		htmlTag:  "p",
		expected: "Main paragraph.",
	},
	{
		name: "p without main",
		input: `<html><body>
		<p>Outside paragraph.</p>
	</body></html>`,
		htmlTag:  "p",
		expected: "Outside paragraph.",
	},
	{
		name:     "no p",
		input:    "<html><body><h1>Test Title</h1></body></html>",
		htmlTag:  "p",
		expected: "",
	},
	{
		name:     "no h1 or h2",
		input:    "<html><body></body></html>",
		htmlTag:  "h1",
		expected: "",
	},
	{
		name:     "no body, get p",
		input:    "",
		htmlTag:  "p",
		expected: "",
	},
	{
		name:     "no body, get h1, h2",
		input:    "",
		htmlTag:  "h1",
		expected: "",
	},
	{
		name:     "get h2, no h1",
		input:    "<html><body><h2>Test Title</h2></body></html>",
		htmlTag:  "h2",
		expected: "Test Title",
	},
}

func TestGetContent(t *testing.T) {
	t.Parallel()

	for _, tt := range tests {
		var got string

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			switch tt.htmlTag {
			case "h1":
				fallthrough
			case "h2":
				got = getHeadingFromHTML(tt.input)
			case "p":
				got = getFirstParagraphFromHTML(tt.input)
			}

			if got != tt.expected {
				t.Errorf("got %q, want %q", got, tt.expected)
			}
		})
	}
}

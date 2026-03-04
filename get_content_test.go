package main

import (
	"testing"
)

var testCases = []struct {
	name     string
	input    string
	htmlTag  string
	expected string
}{
	{
		name:     "GetHeadingFromHTMLBasic",
		input:    "<html><body><h1>Test Title</h1></body></html>",
		htmlTag:  "h1",
		expected: "Test Title",
	},
	{
		name: "GetParagraphFromHTMLMainProperty",
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
		name: "GetParagraphFromHTMLParagraph",
		input: `<html><body>
		<p>Outside paragraph.</p>
	</body></html>`,
		htmlTag:  "p",
		expected: "Outside paragraph.",
	},
	{
		name:     "NoParagraph",
		input:    "<html><body><h1>Test Title</h1></body></html>",
		htmlTag:  "p",
		expected: "",
	},
	{
		name:     "NoHeading",
		input:    "<html><body></body></html>",
		htmlTag:  "h1",
		expected: "",
	},
	{
		name:     "NoBodyGetParagraph",
		input:    "",
		htmlTag:  "p",
		expected: "",
	},
	{
		name:     "NoBodyGetHeading",
		input:    "",
		htmlTag:  "h1",
		expected: "",
	},
	{
		name:     "NoH1GetH2",
		input:    "<html><body><h2>Test Title</h2></body></html>",
		htmlTag:  "h2",
		expected: "Test Title",
	},
}

func TestGetContent(t *testing.T) {
	t.Parallel()

	for _, tc := range testCases {
		var got string

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			switch tc.htmlTag {
			case "h1":
				fallthrough
			case "h2":
				got = getHeadingFromHTML(tc.input)
			case "p":
				got = getFirstParagraphFromHTML(tc.input)
			}

			if got != tc.expected {
				t.Errorf("got %q, want %q", got, tc.expected)
			}
		})
	}
}

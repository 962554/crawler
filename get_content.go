package main

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	H1 = "<h1>"
	H2 = "<h2>"
	P  = "<p>"
)

func getHeadingFromHTML(html string) string {
	if !strings.Contains(html, H1) && !strings.Contains(html, H2) {
		return ""
	}
	buf := strings.NewReader(html)
	doc, err := goquery.NewDocumentFromReader(buf)
	if err != nil {
		return ""
	}
	h1 := doc.Find("h1, h2")
	return h1.First().Text()
}

func getFirstParagraphFromHTML(html string) string {
	if !strings.Contains(html, P) {
		return ""
	}
	buf := strings.NewReader(html)
	doc, err := goquery.NewDocumentFromReader(buf)
	if err != nil {
		return ""
	}

	main := doc.Find("main")
	var p string
	if main.Length() > 0 {
		p = main.Find("p").First().Text()
	} else {
		p = doc.Find("p").First().Text()
	}
	return p
}

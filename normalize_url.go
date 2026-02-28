package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(rawURL string) (string, error) {
	url, err := url.Parse(rawURL)
	if err != nil {
		return "", fmt.Errorf("normalizeURL: failed to parse %s: %w", rawURL, err)
	}

	cleanPath := url.Hostname() + url.EscapedPath()
	cleanPath = strings.TrimRight(cleanPath, "/")
	cleanPath = strings.ToLower(cleanPath)

	return cleanPath, nil
}

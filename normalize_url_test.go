package main

import "testing"

func TestNormalizeURL(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		inputURL string
		expected string
	}{
		{
			name:     "remove https scheme",
			inputURL: "https://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "remove http scheme",
			inputURL: "http://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "remove trailing /",
			inputURL: "https://blog.boot.dev/path/",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "remove miXeD CasE",
			inputURL: "https://blOg.BooT.dEV/paTH/",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "empty string",
			inputURL: "",
			expected: "",
		},
		// add more test cases here
	}

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			actual, err := normalizeURL(tt.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tt.name, err)

				return
			}

			if actual != tt.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tt.name, tt.expected, actual)
			}
		})
	}
}

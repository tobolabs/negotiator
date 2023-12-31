package negotiator_test

import (
	"github.com/noelukwa/negotiator"
	"net/http/httptest"
	"testing"
)

func TestNegotiator_ParseMediaTypes(t *testing.T) {

	cases := []struct {
		name      string
		header    string
		available []string
		expected  []string
	}{
		{
			name:     "should return */*",
			header:   "",
			expected: []string{"*/*"},
		},
		{
			name:   "should return */*",
			header: "text;plain",
			expected: []string{
				"shit",
				"shito",
				"shiti",
			},
		},
		{
			name:   "should return text/*",
			header: "text/*, text/plain;q=0",
			expected: []string{
				"text/*",
			},
		},
		{
			name:   "should return application/json",
			header: "application/json",
			expected: []string{
				"application/json",
			},
		},
		{
			name:   "should return application/json",
			header: "application/json;q=0.5",
			expected: []string{
				"application/json",
			},
		},
		{
			name:     "should return empty slice",
			header:   "application/json;q=0",
			expected: []string{},
		},
		{
			name:   "should return text/html, application/json",
			header: "application/json;q=0.2, text/html",
			expected: []string{
				"text/html",
				"application/json",
			},
		},
		{
			name:   "should return text/*",
			header: "text/*",
			expected: []string{
				"text/*",
			},
		},
		{
			name:   "should return text/plain, text/html, application/json, */*",
			header: "text/plain, application/json;q=0.5, text/html, */*;q=0.1",
			expected: []string{
				"text/plain",
				"text/html",
				"application/json",
				"*/*",
			},
		},
		{
			name:   "should return preferred in order",
			header: "text/plain, application/json;q=0.5, text/html, text/xml, text/yaml, text/javascript, text/csv, text/css, text/rtf, text/markdown, application/octet-stream;q=0.2, */*;q=0.1",
			expected: []string{
				"text/plain",
				"text/html",
				"text/xml",
				"text/yaml",
				"text/javascript",
				"text/csv",
				"text/css",
				"text/rtf",
				"text/markdown",
				"application/json",
				"application/octet-stream",
				"*/*",
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/", nil)
			req.Header.Set("Accept", c.header)

			neg := negotiator.New(req)
			actual := neg.ParseMediaTypes()
			if len(c.expected) >= 1 {
				for i, v := range actual {
					if v != c.expected[i] {
						t.Errorf("Expected %s media type, got %s", c.expected[i], v)
					}
				}
			} else {
				if len(actual) != len(c.expected) {
					t.Errorf("Expected %s media types, got %s", c.expected, actual)
				}
			}

		})
	}

}

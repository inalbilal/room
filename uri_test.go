package room

import (
	"testing"
)

func TestURI_Authority(t *testing.T) {
	uri := URI{authority: "example.com"}
	if uri.Authority() != "example.com" {
		t.Errorf("Authority() returned %s, expected %s", uri.Authority(), "example.com")
	}
}

func TestURI_Path(t *testing.T) {
	uri := URI{path: "/path/to/resource"}
	if uri.Path() != "/path/to/resource" {
		t.Errorf("Path() returned %s, expected %s", uri.Path(), "/path/to/resource")
	}
}

func TestURI_Query(t *testing.T) {
	uri := URI{query: "query=value"}
	if uri.Query() != "query=value" {
		t.Errorf("Query() returned %s, expected %s", uri.Query(), "query=value")
	}
}

func TestURI_Scheme(t *testing.T) {
	uri := URI{scheme: "https"}
	if uri.Scheme() != "https" {
		t.Errorf("Scheme() returned %s, expected %s", uri.Scheme(), "https")
	}
}

func TestURI_String(t *testing.T) {
	uri := URI{
		scheme:    "https",
		authority: "example.com",
		path:      "/path/to/resource",
		query:     "query=value",
	}
	expected := "https://example.com/path/to/resource?query=value"

	if uri.String() != expected {
		t.Errorf("String() returned %s, expected %s", uri.String(), expected)
	}

}

func TestNewURI(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected URI
	}{
		{
			name:  "Simple URL",
			input: "http://example.com/path/to/resource",
			expected: URI{
				scheme:    "http",
				authority: "example.com",
				path:      "/path/to/resource",
				query:     "",
			},
		},
		{
			name:  "URL with Query",
			input: "https://example.com/path/to/resource?query=value",
			expected: URI{
				scheme:    "https",
				authority: "example.com",
				path:      "/path/to/resource",
				query:     "query=value",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uri := NewURI(tt.input)
			if uri.scheme != tt.expected.scheme {
				t.Errorf("Scheme: got %s, want %s", uri.scheme, tt.expected.scheme)
			}
			if uri.authority != tt.expected.authority {
				t.Errorf("Authority: got %s, want %s", uri.authority, tt.expected.authority)
			}
			if uri.path != tt.expected.path {
				t.Errorf("Path: got %s, want %s", uri.path, tt.expected.path)
			}
			if uri.query != tt.expected.query {
				t.Errorf("Query: got %s, want %s", uri.query, tt.expected.query)
			}
		})
	}
}

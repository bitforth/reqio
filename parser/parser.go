package parser

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func ParseFile(filename string) (*http.Request, error) {
    content, err := os.ReadFile(filename)
    if err != nil {
        return nil, err
    }

    // Split the input content into sections, first by double newlines to separate headers from body
    sections := strings.SplitN(string(content), "\n\n", 2)
    // The first part (before a double newline) should contain the start line and headers
    startAndHeaders := strings.Split(sections[0], "\n")
    // The first line should contain at least the HTTP method and the URI
    firstLineParts := strings.SplitN(startAndHeaders[0], " ", 3)
    if len(firstLineParts) < 2 { // Check for method and URI
        return nil, fmt.Errorf("the request must include at least a method and a URI")
    }

    method, uri := firstLineParts[0], firstLineParts[1]
    // Assume there's no body by default
    var body []byte
    if len(sections) > 1 {
        // If there are at least two sections, the second one is considered the body
        body = []byte(sections[1])
    }

    // Create a new HTTP request with the method, URI, and body (if present)
    req, err := http.NewRequest(method, uri, bytes.NewReader(body))
    if err != nil {
        return nil, err
    }

    // Process optional headers if any (skip the first line which is the start line)
    for _, line := range startAndHeaders[1:] {
        if strings.HasPrefix(line, "###") {
            // This is a comment line; ignore
            continue
        }
        headerParts := strings.SplitN(line, ": ", 2)
        if len(headerParts) == 2 {
            req.Header.Set(headerParts[0], headerParts[1])
        }
    }

    return req, nil
}
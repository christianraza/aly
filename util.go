package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func generateURL(query string) string {
	var sb strings.Builder
	sb.WriteString("https://www.google.com")
	_, err := url.ParseRequestURI(query)
	if err != nil {
		sb.WriteString("/search?q=")
		sb.WriteString(url.QueryEscape(query))
	} else {
		sb.WriteString(query)
	}
	return sb.String()
}

func makeRequest(url string) *http.Response {
	client := &http.Client{}

	// Instantiate request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error instantiating request")
		os.Exit(0)
	}
	req.Header.Set("User-Agent", userAgent)

	// Request HTML page
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error requesting HTML page")
		os.Exit(0)
	}

	// Check status code
	if res.StatusCode != 200 {
		fmt.Printf("Status code error: %d, %s\n", res.StatusCode, res.Status)
		os.Exit(0)
	}

	return res
}

func makeDoc(r io.Reader) *goquery.Document {
	// Load HTML document
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		fmt.Println("Error loading HTML document")
		os.Exit(0)
	}
	return doc
}

func checkCorrections(doc *goquery.Document) *goquery.Document {
	n, nURL, o, oURL := parseCorrections(doc)
	if n != "" {
		if o != "" {
			// Use "Showing results for" unless bypassed
			if !bypass {
				fmt.Printf("Searching instead for: %s\n", n)
			} else {
				body := makeRequest(generateURL(oURL)).Body
				doc = makeDoc(body)
			}
		} else {
			// Use "Did you mean" unless bypassed
			if !bypass {
				fmt.Printf("Searching instead for: %s\n", n)
				body := makeRequest(generateURL(nURL)).Body
				doc = makeDoc(body)
			}
		}
	}
	return doc
}

func finalizeOutput(s string) {
	if s != "" {
		fmt.Printf("\n%s\n", strings.TrimSpace(s))
		os.Exit(0)
	}
}

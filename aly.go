package main

import "strings"

const (
	userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:78.0) Gecko/20100101 Firefox/78.0"
)

func main() {
	parseArgs()
	url := generateURL(strings.Join(positionals, " "))
	body := makeRequest(url).Body
	doc := checkCorrections(makeDoc(body))
	finalizeOutput(parseHtml(doc))
}

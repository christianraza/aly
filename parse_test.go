package main

import (
	"testing"
)

func TestParseCorrectionsDidYouMean(t *testing.T) {
	doc := makeDoc(makeRequest(generateURL("water shed")).Body)
	n, nURL, _, _ := parseCorrections(doc)
	if n == "" {
		t.Fatalf("Did you mean text parse failing")
	}
	if nURL == "" {
		t.Fatalf("Did you mean URL parse failing")
	}
}

func TestParseCorrectionsSearchInsteadFor(t *testing.T) {
	doc := makeDoc(makeRequest(generateURL("the mooon is beutiful")).Body)
	_, _, o, oURL := parseCorrections(doc)
	if o == "" {
		t.Fatalf("Search instead for text parse failing")
	}
	if oURL == "" {
		t.Fatalf("Search instead for URL parse failing")
	}
}

func TestParseMath(t *testing.T) {
	doc := makeDoc(makeRequest(generateURL("2 + 2")).Body)
	if parseMath(doc) == "" {
		t.Fatalf("Math parse failing")
	}
}

func TestParseCarousel(t *testing.T) {
	doc := makeDoc(makeRequest(generateURL("cast of lost")).Body)
	if parseCarousel(doc) == "" {
		t.Fatalf("Carousel parse failing")
	}
}

func TestParseBasic(t *testing.T) {
	doc := makeDoc(makeRequest(generateURL("new year's eve")).Body)
	if parseBasic(doc) == "" {
		t.Fatalf("Basic answers parse failing")
	}
}

func TestParseRich(t *testing.T) {
	doc := makeDoc(makeRequest(generateURL("capital of new zealand")).Body)
	if parseRich(doc) == "" {
		t.Fatalf("Rich answers parse failing")
	}
}

func TestParseFeatured(t *testing.T) {
	doc := makeDoc(makeRequest(generateURL("agnes martin bio")).Body)
	if parseFeatured(doc) == "" {
		t.Fatalf("Featured parse failing")
	}
}

func TestParseLyrics(t *testing.T) {
	doc := makeDoc(makeRequest(generateURL("this is a test lyrics")).Body)
	if parseLyrics(doc) == "" {
		t.Fatalf("Lyrics parse failing")
	}
}

func TestParseWeather(t *testing.T) {
	doc := makeDoc(makeRequest(generateURL("weather wisconsin")).Body)
	if parseWeather(doc) == "" {
		t.Fatalf("Weather parse failing")
	}
}

func TestParseUnitConversion(t *testing.T) {
	doc := makeDoc(makeRequest(generateURL("1 m to cm")).Body)
	if parseUnitConversion(doc) == "" {
		t.Fatalf("Unit conversion parse failing")
	}
}

func TestParseCurrencyConversion(t *testing.T) {
	doc := makeDoc(makeRequest(generateURL("1 usd to yen")).Body)
	if parseCurrencyConversion(doc) == "" {
		t.Fatalf("Currency conversion parse failing")
	}
}

func TestParseTranslation(t *testing.T) {
	doc := makeDoc(makeRequest(generateURL("hello in japanese")).Body)
	if parseTranslation(doc) == "" {
		t.Fatalf("Translation parse failing")
	}
}

func TestParseKnowledgePanel(t *testing.T) {
	doc := makeDoc(makeRequest(generateURL("rust-lang")).Body)
	if parseKnowledgePanel(doc) == "" {
		t.Fatalf("Knowledge panel parse failing")
	}
}

func TestParseDefinitions(t *testing.T) {
	doc := makeDoc(makeRequest(generateURL("run def")).Body)
	if parseDefinitions(doc) == "" {
		t.Fatalf("Definitions parse failing")
	}
}

func TestParsePronounce(t *testing.T) {
	doc := makeDoc(makeRequest(generateURL("how to pronounce worcestershire")).Body)
	if parsePronounce(doc) == "" {
		t.Fatalf("Pronounce parse failing")
	}
}

package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/PuerkitoBio/goquery"
)

var (
	bypass      bool
	positionals []string
)

func parseArgs() {
	// Define usage
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "\nUsage\n    %s\n", "aly [flags] <query>")
		fmt.Fprintf(os.Stderr, "Flags\n")
		flag.VisitAll(func(f *flag.Flag) {
			fmt.Fprintf(os.Stderr, "    -%s    %s\n", f.Name, f.Usage)
		})
	}
	// Define flags
	flag.BoolVar(&bypass, "b", false, "bypass the 'Searching instead for' corrections")
	// Parse flags
	flag.Parse()
	// Get positional arguments
	positionals = flag.Args()
	// Require positional arguments
	if flag.NArg() == 0 {
		if flag.NFlag() == 0 {
			fmt.Fprintf(os.Stderr, "\n%s\n", "I'm Aly your personal assistant!")
		}
		flag.Usage()
		os.Exit(0)
	}
}

func parseCorrections(doc *goquery.Document) (n, nURL, o, oURL string) {
	s := doc.Find("#taw")

	s0 := s.Find("a.gL9Hy")
	n = s0.Text()
	nURL, _ = s0.Attr("href")

	s1 := s.Find("a.spell_orig")
	o = s1.Text()
	oURL, _ = s1.Attr("href")
	return
}

func parseMath(doc *goquery.Document) string {
	var b strings.Builder
	doc.Find(".qv3Wpe").Each(func(i int, s *goquery.Selection) {
		b.WriteString(s.Text())
		b.WriteString("\n")
	})
	return b.String()
}

func parseCarousel(doc *goquery.Document) string {
	var b strings.Builder
	doc.Find(".Z8r5Gb.PZPZlf").Each(func(i int, s *goquery.Selection) {
		// Main text
		main := s.Find(".JjtOHd").Text()
		if main != "" {
			b.WriteString(main)
			// Sub text
			sub := s.Find(".ellip.yF4Rkc.AqEFvb").Text()
			if sub != "" {
				b.WriteString(" (")
				b.WriteString(sub)
				b.WriteString(")")
			}
			b.WriteString("\n")
		}
	})
	return b.String()
}

func parseBasic(doc *goquery.Document) string {
	var b strings.Builder
	doc.Find(".zCubwf").Each(func(i int, s *goquery.Selection) {
		b.WriteString(s.Text())
		b.WriteString("\n")
	})
	return b.String()
}

func parseRich(doc *goquery.Document) string {
	var b strings.Builder
	doc.Find(".Z0LcW.XcVN5d").Each(func(i int, s *goquery.Selection) {
		b.WriteString(s.Text())
		b.WriteString("\n")
	})
	return b.String()
}

func parseFeatured(doc *goquery.Document) string {
	var b strings.Builder
	doc.Find(".ILfuVd span").Each(func(i int, s *goquery.Selection) {
		b.WriteString(s.Text())
		b.WriteString("\n")
	})
	return b.String()
}

func parseLyrics(doc *goquery.Document) string {
	var b strings.Builder
	doc.Find("span[jsname=\"YS01Ge\"]").Each(func(i int, s *goquery.Selection) {
		b.WriteString(s.Text())
		b.WriteString("\n")
	})
	return b.String()
}

func parseWeather(doc *goquery.Document) string {
	var b strings.Builder
	doc.Find(".nawv0d").Each(func(i int, s *goquery.Selection) {
		// Cloud conditions
		v, _ := s.Find("#wob_tci").Attr("alt")
		b.WriteString(v)
		b.WriteString("\n")
		// Temperature in degrees
		b.WriteString(s.Find("#wob_tm").Text())
		b.WriteString(s.Find(".wob-unit > span").First().Text())
		b.WriteString("\n")
		// General weather statistics
		s.Find(".wtsRwe div").Each(func(i int, s *goquery.Selection) {
			b.WriteString(s.Contents().Not("span").Text())
			b.WriteString(s.Find("span").Contents().First().Text())
			b.WriteString("\n")
		})
	})
	return b.String()
}

func parseUnitConversion(doc *goquery.Document) string {
	var b strings.Builder
	doc.Find("[jsname=\"fPLMtf\"]").Each(func(i int, s *goquery.Selection) {
		v, _ := s.Attr("value")
		b.WriteString(v)
		b.WriteString("\n")
	})
	return b.String()
}

func parseCurrencyConversion(doc *goquery.Document) string {
	var b strings.Builder
	doc.Find(".SwHCTb").Each(func(i int, s *goquery.Selection) {
		b.WriteString(s.Text())
		b.WriteString("\n")
	})
	return b.String()
}

func parseTranslation(doc *goquery.Document) string {
	var b strings.Builder
	doc.Find("#tw-ob").Each(func(i int, s *goquery.Selection) {
		b.WriteString(s.Find("#tw-source-text-container span").Text())
		b.WriteString(" \u2B82 ")
		b.WriteString(s.Find("#tw-target-text-container span").Text())
		b.WriteString(" (")
		b.WriteString(s.Find("#tw-target-rmn-container span").Text())
		b.WriteString(")\n")
	})
	return b.String()
}

func parseKnowledgePanel(doc *goquery.Document) string {
	var b strings.Builder
	s := doc.Find(".kno-rdesc span").First()
	b.WriteString(s.Text())
	if b.String() != "" {
		b.WriteString("\n")
	}
	return b.String()
}

func parseDefinitions(doc *goquery.Document) string {
	var b strings.Builder
	doc.Find("[jsname=\"r5Nvmf\"]").Each(func(i int, s *goquery.Selection) {
		// Part of speech
		t := s.Find(".pgRvse.vdBwhd.ePtbIe span").First().Text()
		if t != "" {
			b.WriteString(t)
			b.WriteString("\n")
			n := utf8.RuneCountInString(t)
			b.WriteString(strings.Repeat("-", n))
			b.WriteString("\n")
		}
		polyseme := 0
		s.Find("[jsname=\"gskXhf\"]").Each(func(i int, s *goquery.Selection) {
			// Polyseme
			s.Find(".L1jWkf.h3TRxf").Each(func(i int, s *goquery.Selection) {
				polyseme += 1
				b.WriteString(strconv.Itoa(polyseme))
				b.WriteString(". ")
				// Meaning
				b.WriteString(s.Find("[data-dobid=\"dfn\"]").Text())
				b.WriteString("\n")
				// Usage
				t := s.Find(".H9KYcb").First().Text()
				if t != "" {
					b.WriteString(t)
					b.WriteString("\n")
				}
			})
		})
		b.WriteString("\n")
	})
	return b.String()
}

func parsePronounce(doc *goquery.Document) string {
	var b strings.Builder
	doc.Find(".Jzw6hb span").Each(func(i int, s *goquery.Selection) {
		b.WriteString(s.Text())
	})
	if b.String() != "" {
		b.WriteString("\n")
	}
	return b.String()
}

func parseHtml(doc *goquery.Document) string {
	if s := parseMath(doc); s != "" {
		return s
	}

	if s := parseCarousel(doc); s != "" {
		return s
	}

	if s := parseBasic(doc); s != "" {
		return s
	}

	if s := parseRich(doc); s != "" {
		return s
	}

	if s := parseFeatured(doc); s != "" {
		return s
	}

	if s := parseLyrics(doc); s != "" {
		return s
	}

	if s := parseWeather(doc); s != "" {
		return s
	}

	if s := parseUnitConversion(doc); s != "" {
		return s
	}

	if s := parseCurrencyConversion(doc); s != "" {
		return s
	}

	if s := parseTranslation(doc); s != "" {
		return s
	}

	if s := parseKnowledgePanel(doc); s != "" {
		return s
	}

	if s := parseDefinitions(doc); s != "" {
		return s
	}

	if s := parsePronounce(doc); s != "" {
		return s
	}

	return "No results found"
}

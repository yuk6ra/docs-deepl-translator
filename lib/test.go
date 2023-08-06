package lib

import (
	// "net/http"
	"fmt"
	// "regexp"
	// "io"
	"bytes"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"


	"github.com/PuerkitoBio/goquery"

	"strings"
)

func Test() {

	// a, _ := http.Get("https://github.com/russross/blackfriday/blob/master/README.md?plain=1")
	// defer a.Body.Close()

	// body, _ := io.ReadAll(a.Body)

	// fmt.Println(
	// 	string(body),
	// )

	md := `
# High-level Overview

## What is the Cosmos SDK

The [Cosmos SDK](https://github.com/cosmos/cosmos-sdk) is an open-source framework for building multi-asset public Proof-of-Stake (PoS) blockchains, like the Cosmos Hub, as well as permissioned Proof-of-Authority (PoA) blockchains. Blockchains built with the Cosmos SDK are generally referred to as application-specific blockchains.
`

	// Parse Markdown and extract the English part
	englishPart := extractEnglish(md)


	// Translate the English part to Japanese
	translatedText := translateToJapanese(englishPart)

	fmt.Println(translatedText)

	// Replace the English part with the translated text in the original Markdown
	translatedMd := strings.Replace(md, englishPart, translatedText, 1)

	fmt.Println(translatedMd)

}

// Extract the English part from the Markdown using Goldmark
func extractEnglish(md string) string {
	var englishPart string

	mdParser := goldmark.New(
		goldmark.WithExtensions(extension.DefinitionList),
		goldmark.WithParserOptions(
			parser.WithBlockParsers(),
			// parser.WithIDs(),
			parser.WithAttribute(),
			// parser.WithAutoHeadingID(),
		),
	)


	source := []byte(md)

	var buf strings.Builder
	if err := mdParser.Convert(source, &buf); err != nil {
		panic(err)
	}

	// parsedMd := buf.String()

	fmt.Println("B",buf.String())

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader([]byte(buf.String())))
	if err != nil {
		panic(err)
	}

	doc.Find("*").Each(func(i int, s *goquery.Selection) {
		tag := s.Get(0).Data
		text := s.Text()
		fmt.Printf("%s: %s\n", tag, text)
		// fmt.Printf(tag, text)
	})

	return englishPart
}

// Translate the English text to Japanese (Placeholder function, actual translation code needed)
func translateToJapanese(englishText string) string {
	// Placeholder translation function
	// Replace this with your actual translation logic or API call
	// For the sake of demonstration, let's just add "[日本語訳]" at the beginning of the text
	return "[日本語訳]\n" + englishText
}
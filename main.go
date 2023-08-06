package main

import (
	"deepl/lib"
	"fmt"
	"os"
	"strings"
	"bytes"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"

	"github.com/PuerkitoBio/goquery"
)

func main() {

	str, err := os.ReadFile("sample.md")

	// Translate the English part into Japanese
	
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(str))
	
	md := string(str)


	// var englishPart string

	mdParser := goldmark.New(
		goldmark.WithExtensions(extension.DefinitionList),
		goldmark.WithParserOptions(
			parser.WithAttribute(),
			parser.WithAutoHeadingID(),
		),
	)


	source := []byte(md)

	var buf strings.Builder
	if err := mdParser.Convert(source, &buf); err != nil {
		panic(err)
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader([]byte(buf.String())))
	if err != nil {
		panic(err)
	}

	doc.Find("*").Each(func(i int, s *goquery.Selection) {
		tag := s.Get(0).Data
		text := s.Text()
		fmt.Printf("%s: %s\n", tag, text)
	})

	a := strings.Split(buf.String(), "\n")

	for _, text := range a {
		if strings.HasPrefix(text, "<h") {
			fmt.Println(text)
		} else {
			deepl := lib.DeeplRequest{
				Text:        text,
				Source_lang: "EN",
				Target_lang: "JA",
			}

			fmt.Println(text)
			result := lib.DeepLTransration(deepl)
			fmt.Println(result)
		}
	}

	// lib.Hi()
	// fmt.Println("Hi!")

	// bytes, err := os.ReadFile("sample.md")

	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(string(bytes))

	// d := lib.DeeplRequest{
	// 	Text:        string(bytes),
	// 	Source_lang: "EN",
	// 	Target_lang: "JA",
	// }

	// fmt.Println(d)
	// result := lib.DeepLTransration(d)
	// fmt.Println(result)

	

}
package lib

import (
	"fmt"
	"strings"

	"github.com/JohannesKaufmann/html-to-markdown"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
)

func ConvertHTML(md string) string {

	mdParser := goldmark.New(
		goldmark.WithExtensions(extension.DefinitionList),
		goldmark.WithParserOptions(
			parser.WithBlockParsers(),
			parser.WithAttribute(),
			// parser.WithIDs(),
			// parser.WithAutoHeadingID(),
		),
	)

	source := []byte(md)

	var buf strings.Builder
	if err := mdParser.Convert(source, &buf); err != nil {
		fmt.Println(err)
	}

	return buf.String()
}

func ConvertMarkdown(html string) string {
	converter := md.NewConverter("", true, nil)

	markdown, err := converter.ConvertString(html)
	if err != nil {
		fmt.Println(err)
	}

	return markdown
}

func GetMetadata(md string) string {
	metadata := strings.Split(md, "---")[1]
	return ("---" + metadata + "---")
}

func ReplaceAsterisks(input string) string {
	replaced := strings.ReplaceAll(input, "* ", "- ")
	return replaced
}

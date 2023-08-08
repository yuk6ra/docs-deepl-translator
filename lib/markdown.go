package lib


import (
	"fmt"
	"strings"
	"golang.org/x/net/html"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/JohannesKaufmann/html-to-markdown"

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

func SplitHTMLByTags(htmlContent string, tags ...string) []string {
	var sections []string

	reader := strings.NewReader(htmlContent)
	tokenizer := html.NewTokenizer(reader)

	for {
		tokenType := tokenizer.Next()
		if tokenType == html.ErrorToken {
			break
		}

		token := tokenizer.Token()
		for _, tag := range tags {
			if tokenType == html.StartTagToken && token.Data == tag {
				section := extractSection(tokenizer)
				sections = append(sections, section)
				break
			}
		}
	}

	return sections
}

func extractSection(tokenizer *html.Tokenizer) string {
	var sectionBuilder strings.Builder

	for {
		tokenType := tokenizer.Next()
		if tokenType == html.ErrorToken || (tokenType == html.EndTagToken && tokenizer.Token().Data == "h2") {
			break
		}

		token := tokenizer.Token()
		sectionBuilder.WriteString(token.String())
	}

	return sectionBuilder.String()
}
package main

import (
	// "encoding/json"
	"deepl/lib"
	"fmt"
	"os"
	"regexp"
	"strings"

)

func main() {
	input := "docs/input/overview.md"
	output := "docs/output/overview.md"

	str, err := os.ReadFile(input)
	
	if err != nil {
		fmt.Println(err)
	}

	var mdAll string

	metadata := lib.GetMetadata(string(str))
	
	re := regexp.MustCompile(regexp.QuoteMeta(metadata))

	mdAll = string(str)

	md := lib.ReplaceAsterisks(mdAll)
	fmt.Println("MD", md)

	mdSplit := re.Split(string(md), 2)[1]

	htmlAll := lib.ConvertHTML(mdSplit)

	// htmlSlice := lib.SplitHTMLByTags(htmlAll, "h2", "ul", "h3", "h4", "h5", "h6", "p")
	htmlSlice := strings.Split(htmlAll, "\n")
	fmt.Println("HTML", htmlSlice)

	resultSlice := []string{metadata}

	for i, text := range htmlSlice {
		fmt.Println(i, text)

		switch {
		case strings.Contains(text, "<h"):
			md := lib.ConvertMarkdown(text)
			resultSlice = append(resultSlice, md)
		case strings.Contains(text, "<li>"):
			req := lib.DeeplRequest{
				Text:        text,
				Source_lang: "EN",
				Target_lang: "JA",
			}
			res := lib.DeepLTransration(req)
			md := lib.ConvertMarkdown(res)
			resultSlice = append(resultSlice, "- " + md)
		case strings.Contains(text, ":::"):
			md := lib.ConvertMarkdown(text)
			resultSlice = append(resultSlice, "\n" + md + "\n")
		default:
			req := lib.DeeplRequest{
				Text:        text,
				Source_lang: "EN",
				Target_lang: "JA",
			}
			res := lib.DeepLTransration(req)
			md := lib.ConvertMarkdown(res)

			resultSlice = append(resultSlice, md)			
		}
	}

	fmt.Println("RESULT", resultSlice)

	file, err := os.Create(output)
	defer file.Close()

	if err != nil {
		fmt.Println(err)
	}

	data := []byte(strings.Join(resultSlice, "\n"))
	
	_, err = file.Write(data)
	
	if err != nil {
		fmt.Println(err)
	}

}
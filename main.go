package main

import (
	"deepl/lib"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {

	targetPath := "docs/"
	
	markdownFiles, err := getAllMarkdownFiles(targetPath)

	if err != nil {
		fmt.Println(err)
		return
	}

	for i, mdFile := range markdownFiles {
		if i == 1 {
			fmt.Println(mdFile)
			translated := translate(mdFile)

			outputPath := strings.Replace(mdFile, "docs/", "docs_ja/", 1)

			write(outputPath, translated)
		}
	}

}


func getAllMarkdownFiles(rootPath string) ([]string, error) {
	var mdFiles []string

	err := filepath.WalkDir(rootPath, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && filepath.Ext(d.Name()) == ".md" {
			mdFiles = append(mdFiles, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	fmt.Println(mdFiles)

	return mdFiles, nil
}

func write(filepath string, resultSlice []string) {

	dir := filepath[:strings.LastIndex(filepath, "/")]
	fmt.Println(dir)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}

	file, err := os.Create(filepath)
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


func translate(input string) []string {
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
			resultSlice = append(resultSlice, "- "+md)
		case strings.Contains(text, ":::"):
			md := lib.ConvertMarkdown(text)
			resultSlice = append(resultSlice, "\n"+md+"\n")
		case strings.Contains(text, "<code"):
			md := lib.ConvertMarkdown(text)
			resultSlice = append(resultSlice, "\n"+md+"\n")
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

	return resultSlice
}
package lib

import (
	"fmt"
	"os"

	"github.com/russross/blackfriday"
)

func Mark() {
	fmt.Println("Hi!")

	bytes, err := os.ReadFile("assets/overview.md")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bytes))

	output := blackfriday.MarkdownBasic([]byte(string(bytes)))
	fmt.Println(string(output))
	
}

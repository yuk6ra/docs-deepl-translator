package lib

import (
	"fmt"
	"io"
	"os"
	"net/http"
	"net/url"
	"strings"

	"github.com/joho/godotenv"
)

const END_POINT string = "https://api-free.deepl.com/v2/translate"


type DeeplRequest struct {
	Text string
	Source_lang string
	Target_lang string
}

func DeepLTransration(vals DeeplRequest) string {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env file")
	}

	api_key := os.Getenv("DEEPL_API_KEY")

	data := url.Values{}
	data.Set("text", vals.Text)
	data.Set("source_lang", vals.Source_lang)
	data.Set("target_lang", vals.Target_lang)

	req, err := http.NewRequest(
		"POST",
		END_POINT,
		nil,
	)

	if err != nil {
		fmt.Println("error", err)
	}

	req.Header.Set("Authorization", "DeepL-Auth-Key " + api_key)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	req.Body = io.NopCloser(strings.NewReader(data.Encode()))

	client := &http.Client{}
	resp, err := client.Do(req)
	
	if err != nil {
		fmt.Println("error", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("error", err)
	}

	fmt.Println("Response:", string(body))

	return string(body)
}
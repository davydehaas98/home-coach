package util

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

func CreateGetRequest(url string) *http.Request {
	return createRequest("GET", url, nil)
}

func CreatePostRequest(url string, body *url.Values) *http.Request {
	return createRequest("POST", url, []byte(body.Encode()))
}

func createRequest(method string, url string, body []byte) *http.Request {
	request, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		log.Fatal("Could not create request.", err)
	}
	return request
}

func DoRequest(request *http.Request) *http.Response {
	client := http.Client{Timeout: 10 * time.Second}
	response, err := client.Do(request)
	if err != nil {
		log.Fatal("Could not send request.", err)
	}
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		log.Fatal("Request failed with status code: ", response.StatusCode)
	}

	return response
}

func UnmarshalJson[T any](response *http.Response) T {
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	log.Println(string(responseBody))
	if err != nil {
		log.Fatal("Could not get response body.", err)
	}

	var result T
	if err := json.Unmarshal(responseBody, &result); err != nil {
		log.Fatal("Could not unmarshal response body.")
	}

	return result
}

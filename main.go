package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

const BASE_URL = "https://api.netatmo.com"
const TOKEN_URL = BASE_URL + "/oauth2/token"

func main() {
	token := refreshToken()
	log.Println(token)

	log.Println("Starting Ticker..")
	ticker := time.NewTicker(5 * time.Second)

	for t := range ticker.C {
		log.Println("Ticked at", t)
	}

	ticker.Stop()
}

func getEnv(key string) string {
	env := os.Getenv(key)
	if env == "" {
		log.Fatalf("%s not set", key)
	}
	return env
}

func refreshToken() string {
	body := url.Values{}
	body.Set("client_id", getEnv("HC_CLIENT_ID"))
	body.Set("client_secret", getEnv("HC_CLIENT_SECRET"))
	body.Set("grant_type", "refresh_token")
	body.Set("refresh_token", getEnv("HC_REFRESH_TOKEN"))

	log.Println("Refreshing token..")
	request, err := http.NewRequest("POST", TOKEN_URL, bytes.NewReader([]byte(body.Encode())))
	if err != nil {
		log.Fatal("Could not create request.", err)
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	log.Println(request)

	client := http.Client{Timeout: 10 * time.Second}
	response, err := client.Do(request)
	if err != nil {
		log.Fatal("Could not send request.", err)
	}
	if response.StatusCode != 200 {
		log.Fatal("Could not refresh token.", response.Status)
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Could not get response body.", err)
	}

	log.Println("Refreshed token.", string(responseBody))
	return string(responseBody)
}

type RefreshTokenPostRequest struct {
	ClientId     string `form:"client_id"`
	ClientSecret string `form:"client_secret"`
	GrantType    string `form:"grant_type"`
	RefreshToken string `form:"refresh_token"`
}

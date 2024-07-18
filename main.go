package main

import (
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

func refreshToken() string {
	clientId := os.Getenv("HC_CLIENT_ID")
	if clientId == "" {
		// log.Fatal(("HC_CLIENT_ID not set"))
	}
	clientSecret := os.Getenv("HC_CLIENT_ID")
	if clientSecret == "" {
		// log.Fatal(("HC_CLIENT_SECRET not set"))
	}
	grantType := "refresh_token"
	refreshToken := os.Getenv("HC_REFRESH_TOKEN")
	// expiration := 0

	url, err := url.Parse(TOKEN_URL)
	if err != nil {
		// log.Fatal("TOKEN_URL could not be parsed", err)
	}

	query := url.Query()
	query.Set("grant_type", grantType)
	query.Set("refresh_token", refreshToken)
	query.Set("client_id", clientId)
	query.Set("client_secret", clientSecret)
	url.RawQuery = query.Encode()

	log.Println("Refreshing token..", url.RawQuery)
	response, err := http.Post(url.String(), "application/json", nil)
	if err != nil {
		log.Fatal("Could not refresh token.", err)
	}
	log.Println("Refreshed token.", response)
	defer func(body io.ReadCloser) {
		_ = body.Close()
	}(response.Body)

	token, _ := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Could not get token.", err)
	}

	return string(token)
}

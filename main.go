package main

import (
	"home-coach/model"
	"home-coach/util"
	"log"
	"net/url"
	"time"
)

const BASE_URL = "https://api.netatmo.com"
const TOKEN_URL = BASE_URL + "/oauth2/token"
const DATA_URL = BASE_URL + "/api/gethomecoachsdata"

var env map[string]string

func init() {
	env = util.LoadEnv()
}

func main() {
	log.Println("Starting ticker..")
	ticker := time.NewTicker(5 * time.Second)

	for t := range ticker.C {
		if util.IsExpired(env["HC_EXPIRATION"]) {
			refreshAccessToken()
		}
		getData()
		log.Println("Ticked at", t)
	}

	ticker.Stop()
}

func getData() model.DataResponse {
	log.Println("Retrieving data..")

	request := util.CreateGetRequest(DATA_URL)
	bearer := "Bearer " + env["HC_ACCESS_TOKEN"]
	request.Header.Set("Authorization", bearer)
	request.Header.Set("accept", "application/json")

	response := util.DoRequest(request)
	log.Println(response)

	result := util.UnmarshalJson[model.DataResponse](response)

	log.Println("Retrieved data.")

	return result
}

func refreshAccessToken() {
	log.Println("Refreshing access token..")

	body := url.Values{}
	body.Set("client_id", env["HC_CLIENT_ID"])
	body.Set("client_secret", env["HC_CLIENT_SECRET"])
	body.Set("grant_type", "refresh_token")
	body.Set("refresh_token", env["HC_REFRESH_TOKEN"])

	request := util.CreatePostRequest(TOKEN_URL, &body)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response := util.DoRequest(request)

	result := util.UnmarshalJson[model.RefreshTokenResponse](response)

	// Set new tokens in .env file
	env = util.SetEnv("HC_REFRESH_TOKEN", result.RefreshToken)
	env = util.SetEnv("HC_ACCESS_TOKEN", result.AccessToken)
	expiration := time.Now().UTC().Add(time.Second * time.Duration(result.ExpiresIn)).Format(time.RFC3339)
	env = util.SetEnv("HC_EXPIRATION", expiration)

	log.Println("Refreshed token.")
}

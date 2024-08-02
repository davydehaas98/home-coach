package util

import (
	"log"
	"time"

	"github.com/joho/godotenv"
)

const ENV_FILE = ".env"

func LoadEnv() map[string]string {
	envMap, err := godotenv.Read(ENV_FILE)
	if err != nil {
		log.Fatalf("Error loading '%s' file: %s", ENV_FILE, err)
	}
	return envMap
}

func SetEnv(key, value string) map[string]string {
	env := LoadEnv()
	env[key] = value
	godotenv.Write(env, ENV_FILE)
	// Reload env
	return LoadEnv()
}

func IsExpired(expiration string) bool {
	exp, err := time.Parse(time.RFC3339, expiration)
	if err != nil {
		log.Fatal("Could not parse expiration time.", err)
	}
	return time.Now().After(exp)
}

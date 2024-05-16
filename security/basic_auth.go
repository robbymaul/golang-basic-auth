package security

import (
	"crypto/sha256"
	"crypto/subtle"
	"os"

	"github.com/joho/godotenv"
)

func BasicAuth(u string, p string) bool {
	match := match(u, p)
	if !match {
		return false
	}

	return true
}

func match(username string, password string) bool {
	godotenv.Load("../.env")
	usernameENV := os.Getenv("BASIC_AUTH_USERNAME")
	passwordENV := os.Getenv("BASIC_AUTH_PASSWORD")

	usernameHash := sha256.Sum256([]byte(username))
	passwordHash := sha256.Sum256([]byte(password))

	basicAuthUsername := sha256.Sum256([]byte(usernameENV))
	basicAuthPassword := sha256.Sum256([]byte(passwordENV))

	usernameMatch := (subtle.ConstantTimeCompare(usernameHash[:], basicAuthUsername[:]) == 1)
	passwordMatch := (subtle.ConstantTimeCompare(passwordHash[:], basicAuthPassword[:]) == 1)

	if !usernameMatch || !passwordMatch {
		return false
	}

	return true
}

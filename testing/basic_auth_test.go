package testing

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/robbymaul/golang-basic-auth.git/security"
	"github.com/stretchr/testify/assert"
)

type Auth struct {
	Username string
	Password string
}

func TestBasicSuccess(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		t.Fatal(err.Error())
	}

	auth := Auth{
		Username: os.Getenv("BASIC_AUTH_USERNAME"),
		Password: os.Getenv("BASIC_AUTH_PASSWORD"),
	}

	match := security.BasicAuth(auth.Username, auth.Password)

	assert.Equal(t, true, match)
}

func TestBasicFailed(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		t.Fatal(err.Error())
	}

	auth := Auth{
		Username: "",
		Password: "",
	}

	match := security.BasicAuth(auth.Username, auth.Password)

	assert.Equal(t, false, match)
}

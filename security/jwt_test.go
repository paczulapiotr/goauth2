package security

import (
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func TestCreateAuthoizationCode(t *testing.T) {
	codeBase := "auth code base stringz"
	code, validUntil := CreateAuthorizationCode(codeBase)

	expectedCodeStringLength := 44
	actualLength := len(code)
	if actualLength != expectedCodeStringLength {
		t.Fatalf("Authorization code has invalid lenght")
	}

	if validUntil.Before(time.Now()) {
		t.Fatalf("Auth code valid until date in invalid")
	}
}

func TestCreateAccessToken(t *testing.T) {
	accessTokenSecret := "TestingAccessSecret"
	token, _, err := CreateAccessToken("21345om", "newuser", accessTokenSecret)

	if err != nil {
		t.Fatalf("Could not create token: %v", err.Error())
	}

	jwt, err := jwt.Parse(token, func(tkn *jwt.Token) (interface{}, error) { return []byte(accessTokenSecret), nil })

	if !jwt.Valid {
		t.Fatalf("JWT token is invalid")
	}
}

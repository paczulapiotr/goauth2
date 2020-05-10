package security

import (
	"crypto/sha256"
	"encoding/base64"
	"math"
	"math/rand"
	"strconv"
	"time"
)

var authCodeValidPeriod int = 30

// CreateAuthorizationCode creates new auth code based on login
func CreateAuthorizationCode(login string) (authCode string, validUntil time.Time) {
	rand.Seed(time.Now().UnixNano())
	randInteger := rand.Intn(math.MaxInt16)
	hashBase := login + strconv.Itoa(randInteger)
	hash := sha256.Sum256([]byte(hashBase))

	authCode = base64.StdEncoding.EncodeToString(hash[:])
	validUntil = time.Now().UTC().Add(time.Minute * time.Duration(authCodeValidPeriod))
	return
}

// CreateRefreshToken creates refresh token
func CreateRefreshToken() (refreshToken string, validUntil time.Time, err error) {
	return "REFRESH_TOKEN", time.Now(), nil
}

// CreateAccessToken creates access token
func CreateAccessToken() (accessToken string, validUntil time.Time, err error) {
	return "ACCESS_TOKEN", time.Now(), nil
}

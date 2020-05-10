package security

import (
	"crypto/sha256"
	"encoding/base64"
	"math"
	"math/rand"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/paczulapiotr/goauth2/config"
)

var authCodeValidPeriod int = 30
var cfg = config.GetConfiguration()
var refreshTokenExpirationDuration = time.Hour * 24 * 7
var accessTokenExpirationDuration = time.Minute * 60

// CreateAuthorizationCode creates new auth code based on login
func CreateAuthorizationCode(login string) (authCode string, validUntil time.Time) {
	authCode = createRandomSha256(login)
	validUntil = time.Now().UTC().Add(time.Minute * time.Duration(authCodeValidPeriod))
	return
}

// CreateRefreshToken creates refresh token
func CreateRefreshToken(login string) (refreshToken string, validUntil time.Time, err error) {
	refreshToken = createRandomSha256(login + cfg.RefreshTokenSecret)
	validUntil = time.Now().UTC().Add(refreshTokenExpirationDuration)
	return
}

// CreateAccessToken creates access token
func CreateAccessToken(userID, login string) (accessToken string, validUntil time.Time, err error) {
	now := time.Now().UTC()
	validUntil = now.Add(accessTokenExpirationDuration)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Id:        userID,
		ExpiresAt: validUntil.Unix(),
		Subject:   login,
		Issuer:    login,
		NotBefore: now.Unix(),
		IssuedAt:  now.Unix(),
		Audience:  "all",
	})
	secret := []byte(cfg.AccessTokenSecret)
	accessToken, err = jwtToken.SignedString(secret)

	return
}

func createRandomSha256(key string) string {
	rand.Seed(time.Now().UnixNano())
	randInteger := rand.Intn(math.MaxInt16)
	hashBase := key + strconv.Itoa(randInteger)
	hash := sha256.Sum256([]byte(hashBase))

	return base64.StdEncoding.EncodeToString(hash[:])
}

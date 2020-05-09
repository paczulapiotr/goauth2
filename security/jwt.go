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

// CreateAuthoizationCode creates new auth code based on login
func CreateAuthoizationCode(login string) (authCode string, validUntil time.Time) {
	rand.Seed(time.Now().UnixNano())
	randInteger := rand.Intn(math.MaxInt16)
	hashBase := login + strconv.Itoa(randInteger)
	hash := sha256.Sum256([]byte(hashBase))

	authCode = base64.StdEncoding.EncodeToString(hash[:])
	validUntil = time.Now().Add(time.Minute * time.Duration(authCodeValidPeriod))
	return
}

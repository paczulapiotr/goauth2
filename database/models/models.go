package models

import (
	"time"

	"github.com/paczulapiotr/goauth2/security"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User struct for mapping mongo data
type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Login        string             `bson:"login,omitempty"`
	PasswordHash string             `bson:"passwordHash,omitempty"`
	Auth         OAuth2             `bson:"auth"`
}

// CheckPassword check password for UserDao struct
func (user User) CheckPassword(password string) bool {
	return security.CheckPasswordHash(user.PasswordHash, password)
}

// OAuth2 struct for JWT data
type OAuth2 struct {
	Code              string    `bson:"code"`
	CodeValidUntil    time.Time `bson:"codeValidUntil"`
	AccessToken       string    `bson:"accessToken"`
	ValidUntil        time.Time `bson:"validUntil"`
	RefreshToken      string    `bson:"refreshToken"`
	RefreshValidUntil time.Time `bson:"refreshValidUntil"`
	Claims            []Claim   `bson:"claims"`
	LastAuth          time.Time `bson:"lastAuth"`
}

// Claim struct for JWT claims
type Claim struct {
	Key   string
	Value string
}
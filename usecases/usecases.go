package usecases

import (
	"errors"
	"fmt"
	"time"

	"github.com/paczulapiotr/goauth2/database"
	"github.com/paczulapiotr/goauth2/security"
)

// LoginForAuthorizationCode Returns auth code for given credentials
func LoginForAuthorizationCode(login, password string) (code string, err error) {

	mongo := database.DefaultClient()
	user, err := database.FindUserByLogin(mongo, login)

	if err != nil {
		return
	}

	if !security.CheckPasswordHash(password, user.PasswordHash) {
		err = errors.New("Invalid credentials")
	}

	if err != nil {
		return
	}
	code, validUntil := security.CreateAuthorizationCode(login)
	err = database.UpdateAuthorizationCode(mongo, login, code, validUntil)

	return
}

// RegisterUser registers new user
func RegisterUser(login, password string) (err error) {
	mongo := database.DefaultClient()

	// validate login
	err = security.ValidateLoginStructure(login)

	if err != nil {
		return
	}

	user, _ := database.FindUserByLogin(mongo, login)

	if user.Login == login {
		err = fmt.Errorf("User with login %v already exists", login)
	}

	if err != nil {
		return
	}

	// validate password
	err = security.CheckPasswordStructure(password)

	if err != nil {
		return
	}

	// add user
	_, err = database.AddUser(mongo, login, password)
	return
}

// UseAuthorizationCode exchanges authorization code for access_token and refresh_token
func UseAuthorizationCode(code string) (
	accessToken string,
	validUntil time.Time,
	refreshToken string,
	refreshValidUntil time.Time,
	err error) {
	// validate authorization code format

	mongo := database.DefaultClient()
	// find authorization code
	user, err := database.FindUserByAuthorizationCode(mongo, code)
	if err != nil {
		return
	}
	// check if still valid (date)
	validCode := user.Auth.CodeValidUntil.UTC().After(time.Now().UTC())

	if !validCode {
		err = errors.New("Authorization code is expired")
	}
	// remove authorization_code from db
	database.UpdateAuthorizationCode(mongo, user.Login, "", time.Time{})

	// create access_token, refresh_token with claims + validity dates

	refreshToken, refreshValidUntil, err = security.CreateRefreshToken()
	if err != nil {
		return

	}
	accessToken, validUntil, err = security.CreateAccessToken()
	if err != nil {
		return
	}

	// add access_token, refresh_token to user document
	err = database.UpdateRefreshToken(mongo, user.Login, refreshToken, accessToken, refreshValidUntil, validUntil)

	// return access_token, refresh_token
	return
}

package usecases

import (
	"errors"
	"fmt"

	"github.com/paczulapiotr/goauth2/database"
	"github.com/paczulapiotr/goauth2/security"
)

// LoginForAuthorizationCode Returns auth code for given credentials
func LoginForAuthorizationCode(login, password string) (code string, err error) {

	mongo := database.DefaultClient()
	user, err := database.FindUser(mongo, login)

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
	err = database.CreateAuthorizationCode(mongo, login, code, validUntil)

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

	user, _ := database.FindUser(mongo, login)

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

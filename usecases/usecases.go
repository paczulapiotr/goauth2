package usecases

import (
	"errors"

	"github.com/paczulapiotr/goauth2/config"
	"github.com/paczulapiotr/goauth2/database"
	"github.com/paczulapiotr/goauth2/security"
)

// LoginForAuthorizationCode Returns auth code for given credentials
func LoginForAuthorizationCode(login, password string) (code string) {
	config := config.GetConfiguration()
	mongo, err := database.CreateClient(config.Mongo)
	if err != nil {
		panic(err)
	}

	user, err := database.FindUser(mongo, login)

	if err != nil {
		panic(err)
	}

	if !security.CheckPasswordHash(password, user.PasswordHash) {
		panic(errors.New("Invalid credentials"))
	}

	code, validUntil := security.CreateAuthorizationCode(login)
	err = database.CreateAuthorizationCode(mongo, login, code, validUntil)

	if err != nil {
		panic(err)
	}

	return
}

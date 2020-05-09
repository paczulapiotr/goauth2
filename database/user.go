package database

import (
	"fmt"
	"time"

	"github.com/paczulapiotr/goauth2/database/models"
	"github.com/paczulapiotr/goauth2/security"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func getUsersCollections(mongo *mongo.Client) *mongo.Collection {
	return mongo.Database("goauth").Collection("users")
}

// AddUser adds UserDao document into users collection
func AddUser(mongo *mongo.Client, login, password string) (string, error) {
	users := getUsersCollections(mongo)
	ctx := createContext()
	passwordHash, _ := security.HashPassword(password)

	user := models.User{
		Login:        login,
		PasswordHash: passwordHash,
	}
	res, err := users.InsertOne(*ctx, user)
	closeConnection(mongo, ctx)

	if err != nil {
		return "", err
	}

	id, ok := res.InsertedID.(*primitive.ObjectID)

	if !ok {
		return "", fmt.Errorf("Could not add new user: %v", user)
	}

	return id.String(), err
}

// FindUser returns user with given login
func FindUser(mongo *mongo.Client, login string) (*models.User, error) {
	users := getUsersCollections(mongo)
	ctx := createContext()
	var userToFind models.User
	res := *users.FindOne(*ctx, bson.D{{"login", login}})
	err := res.Decode(&userToFind)
	return &userToFind, err
}

// CreateAuthorizationCode Adds authorization code to user
func CreateAuthorizationCode(mongo *mongo.Client, login string, code string, codeValidUntil time.Time) error {
	users := getUsersCollections(mongo)
	ctx := createContext()

	code, validUntil := security.CreateAuthoizationCode(login)

	auth := models.OAuth2{
		Code:           code,
		CodeValidUntil: validUntil,
		Claims:         []models.Claim{},
	}

	filter := bson.D{{"login", login}}
	update := bson.D{
		{
			"$set",
			bson.D{{"auth", auth}},
		},
	}

	res, err := users.UpdateOne(*ctx, filter, update)

	if err != nil {
		return err
	}

	if res.ModifiedCount != 1 {
		return fmt.Errorf("Could not add auth code to login: %v", login)
	}

	return nil
}

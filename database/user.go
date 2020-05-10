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
	(*ctx).Done()

	if err != nil {
		return "", err
	}

	id, ok := res.InsertedID.(primitive.ObjectID)

	if !ok {
		return "", fmt.Errorf("Could not add new user: %v", user)
	}

	return id.String(), err
}

// FindUserByLogin returns user with given login
func FindUserByLogin(mongo *mongo.Client, login string) (*models.User, error) {
	return findUserWithFilter(mongo, bson.D{{"login", login}})
}

// FindUserByAuthorizationCode returns user with given code
func FindUserByAuthorizationCode(mongo *mongo.Client, code string) (*models.User, error) {
	return findUserWithFilter(mongo, bson.D{
		{
			"auth.code", code,
		},
	})
}

func findUserWithFilter(mongo *mongo.Client, filter bson.D) (*models.User, error) {
	users := getUsersCollections(mongo)
	ctx := createContext()
	var userToFind models.User
	res := *users.FindOne(*ctx, filter)
	(*ctx).Done()
	err := res.Decode(&userToFind)
	return &userToFind, err
}

// UpdateAuthorizationCode updates authorization code for login
func UpdateAuthorizationCode(mongo *mongo.Client, login, code string, codeValidUntil time.Time) error {
	fieldsToUpdate := primitive.D{
		{"auth.code", code},
		{"auth.codeValidUntil", codeValidUntil},
	}

	return updateUser(mongo, login, fieldsToUpdate)
}

// UpdateAccessToken updates access token for login
func UpdateAccessToken(mongo *mongo.Client, login, accessToken string, validUntil time.Time) error {
	fieldsToUpdate := primitive.D{
		{"auth.accessToken", accessToken},
		{"auth.validUntil", validUntil},
	}

	return updateUser(mongo, login, fieldsToUpdate)
}

// UpdateRefreshToken updates refresh and access token for login
func UpdateRefreshToken(mongo *mongo.Client, login, refreshToken, accessToken string, refreshValidUntil, validUntil time.Time) error {
	fieldsToUpdate := primitive.D{
		{"auth.refreshToken", refreshToken},
		{"auth.refreshValidUntil", refreshValidUntil},
		{"auth.accessToken", accessToken},
		{"auth.validUntil", validUntil},
	}

	return updateUser(mongo, login, fieldsToUpdate)
}

func updateUser(mongo *mongo.Client, login string, fields primitive.D) error {
	users := getUsersCollections(mongo)
	ctx := createContext()

	filter := bson.D{{"login", login}}

	update := updateBsonWrapper(fields)

	res, err := users.UpdateOne(*ctx, filter, update)

	(*ctx).Done()

	if err != nil {
		return err
	}

	if res.ModifiedCount != 1 {
		return fmt.Errorf("Could not update auth code for login: %v", login)
	}

	return nil
}

func updateBsonWrapper(update bson.D) bson.D {
	return bson.D{
		{
			"$set",
			update,
		},
	}
}

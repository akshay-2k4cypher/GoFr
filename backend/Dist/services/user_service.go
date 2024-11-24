package services

import (
	"context"
	"errors"
	"gofr-registration/config"
	"gofr-registration/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// UserService handles user-related operations
type UserService struct{}

// RegisterUser inserts a new user into the database
func (s *UserService) RegisterUser(user models.User) error {
	collection := config.DB.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Check if the email already exists
	var existingUser models.User
	err := collection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&existingUser)
	if err == nil {
		return errors.New("email already exists")
	}

	// Insert the user into the database
	_, err = collection.InsertOne(ctx, user)
	return err
}

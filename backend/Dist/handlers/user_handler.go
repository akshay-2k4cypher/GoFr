package handlers

import (
	"gofr-registration/models"
	"gofr-registration/services"
	"net/http"

	"github.com/gofr-framework/gofr"
)

type UserHandler struct {
	service *services.UserService
}

// NewUserHandler initializes a new handler
func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// Register handles user registration requests
func (h *UserHandler) Register(ctx *gofr.Context) (interface{}, error) {
	var user models.User

	// Bind request body to user model
	if err := ctx.Bind(&user); err != nil {
		return nil, ctx.RespondWithError(http.StatusBadRequest, "Invalid request payload")
	}

	// Register the user
	err := h.service.RegisterUser(user)
	if err != nil {
		if err.Error() == "email already exists" {
			return nil, ctx.RespondWithError(http.StatusConflict, "Email already registered")
		}
		return nil, ctx.RespondWithError(http.StatusInternalServerError, "Failed to register user")
	}

	return map[string]string{"message": "User registered successfully"}, nil
}

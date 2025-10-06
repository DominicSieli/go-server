package controllers

import (
	"fmt"
	"net/http"

	"github.com/dominicsieli/go-server/auth"
	"github.com/dominicsieli/go-server/config"
	"github.com/dominicsieli/go-server/crypto"
	"github.com/dominicsieli/go-server/types"
	"github.com/dominicsieli/go-server/utilities"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Controller struct {
	repo types.UserRepo
}

func (controller *Controller) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", controller.LoginController).Methods(http.MethodPost)
	router.HandleFunc("/register", controller.RegisterController).Methods(http.MethodPost)
}

func CreateController(repo types.UserRepo) *Controller {
	return &Controller{repo: repo}
}

func (controller *Controller) LoginController(response http.ResponseWriter, request *http.Request) {
	var payload types.LoginUserPayload

	if err := utilities.ParseJSON(request, &payload); err != nil {
		utilities.WriteError(response, http.StatusBadRequest, err)

		return
	}

	if err := utilities.Validate.Struct(payload); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		utilities.WriteError(response, http.StatusBadRequest, fmt.Errorf("invalid payload %v", validationErrors))

		return
	}

	user, err := controller.repo.GetUserByEmail(payload.Email)

	if err != nil {
		utilities.WriteError(response, http.StatusBadRequest, fmt.Errorf("email not found"))

		return
	}

	if !crypto.ValidatePasswordHash(user.Password, []byte(payload.Password)) {
		utilities.WriteError(response, http.StatusBadRequest, fmt.Errorf("password not found"))

		return
	}

	secret := []byte(config.Envs.JWTSecret)
	token, err := auth.CreateToken(secret, user.ID)

	if err != nil {
		utilities.WriteError(response, http.StatusInternalServerError, err)

		return
	}

	utilities.WriteJSON(response, http.StatusOK, map[string]string{"token": token})

}

func (controller *Controller) RegisterController(response http.ResponseWriter, request *http.Request) {
	var payload types.RegisterUserPayload

	if err := utilities.ParseJSON(request, &payload); err != nil {
		utilities.WriteError(response, http.StatusBadRequest, err)

		return
	}

	if err := utilities.Validate.Struct(payload); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		utilities.WriteError(response, http.StatusBadRequest, fmt.Errorf("invalid payload %v", validationErrors))

		return
	}

	if _, err := controller.repo.GetUserByEmail(payload.Email); err != nil {
		utilities.WriteError(response, http.StatusBadRequest, fmt.Errorf("user with email %s already exists", payload.Email))

		return
	}

	hashedPassword, err := crypto.HashPassword(payload.Password)

	if err != nil {
		utilities.WriteError(response, http.StatusInternalServerError, err)

		return
	}

	err = controller.repo.CreateUser(types.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:	   payload.Email,
		Password:  hashedPassword,
	})

	if err != nil {
		utilities.WriteError(response, http.StatusInternalServerError, err)

		return
	}

	utilities.WriteJSON(response, http.StatusCreated, nil)
}

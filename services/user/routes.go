package user

import (
	"fmt"
	"net/http"

	"github.com/dominicsieli/go-server/services/auth"
	"github.com/dominicsieli/go-server/types"
	"github.com/dominicsieli/go-server/utilities"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (handler *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", handler.handleLogin).Methods("POST")
	router.HandleFunc("/register", handler.handleRegister).Methods("POST")
}

func (handler *Handler) handleLogin(response http.ResponseWriter, request *http.Request) {
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

	user, err := handler.store.GetUserByEmail(payload.Email)

	if err != nil {
		utilities.WriteError(response, http.StatusBadRequest, fmt.Errorf("email not found"))

		return
	}

	if !auth.ValidatePasswordHash(user.Password, []byte(payload.Password)) {
		utilities.WriteError(response, http.StatusBadRequest, fmt.Errorf("password not found"))

		return
	}

}

func (handler *Handler) handleRegister(response http.ResponseWriter, request *http.Request) {
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

	if _, err := handler.store.GetUserByEmail(payload.Email); err != nil {
		utilities.WriteError(response, http.StatusBadRequest, fmt.Errorf("user with email %s already exists", payload.Email))

		return
	}

	hashedPassword, err := auth.HashPassword(payload.Password)

	if err != nil {
		utilities.WriteError(response, http.StatusInternalServerError, err)

		return
	}

	err = handler.store.CreateUser(types.User{
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

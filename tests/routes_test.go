package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dominicsieli/go-server/controllers"
	"github.com/dominicsieli/go-server/types"
	"github.com/gorilla/mux"
)

func TestUserController(test *testing.T) {
	userRepo := &MockUserRepo{}
	controller := controllers.CreateController(userRepo)

	test.Run("should fail if the user payload is invalid", func(test *testing.T) {
		payload := types.RegisterUserPayload{
			FirstName: "Dominic",
			LastName:  "Sieli",
			Email:	   "invalid",
			Password:  "password",
		}

		marshalled, _ := json.Marshal(payload)
		request, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))

		if err != nil {
			test.Fatal(err)
		}

		testRequest := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("register", controller.RegisterController)
		router.ServeHTTP(testRequest, request)

		if testRequest.Code != http.StatusBadRequest {
			test.Errorf("expected status code %d, got %d", http.StatusBadRequest, testRequest.Code)
		}
	})

	test.Run("should correctly register the user", func(test *testing.T) {
		payload := types.RegisterUserPayload{
			FirstName: "Dominic",
			LastName:  "Sieli",
			Email:	   "dominicsieli@gmail.com",
			Password:  "password",
		}

		marshalled, _ := json.Marshal(payload)
		request, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))

		if err != nil {
			test.Fatal(err)
		}

		testRequest := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("register", controller.RegisterController)
		router.ServeHTTP(testRequest, request)

		if testRequest.Code != http.StatusCreated {
			test.Errorf("expected status code %d, got %d", http.StatusCreated, testRequest.Code)
		}
	})
}

type MockUserRepo struct{}

func (mockUserRepo *MockUserRepo) GetUserByEmail(email string) (*types.User, error) {
	return nil, fmt.Errorf("user not found")
}

func (mockUserRepo *MockUserRepo) GetUserByID(id int) (*types.User, error) {
	return nil, nil
}

func (mockUserRepo *MockUserRepo) CreateUser(types.User) error {
	return nil
}

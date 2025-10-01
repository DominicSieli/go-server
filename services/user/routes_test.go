package user

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dominicsieli/go-server-template/services/user"
	"github.com/dominicsieli/go-server-template/types"
	"github.com/gorilla/mux"
)

func TestUserServiceHandlers(test *testing.T) {
	userStore := &mockUserStore{}
	handler := user.NewHandler(userStore)

	test.Run("should fail if the user payload is invalid", func(test *testing.T) {
		payload := types.RegisterUserPayload{
			FirstName: "Dominic",
			LastName: "Sieli",
			Email: "",
			Password: "password",
		}

		marshalled, _ := json.Marshal(payload)
		request, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))

		if err != nil {
			test.Fatal(err)
		}

		testRequest := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("register", handler.handleRegister)
		router.ServeHTTP(testRequest, request)

		if testRequest.Code != http.StatusBadRequest {
			test.Errorf("expected status code %d, got %d", http.StatusBadRequest, testRequest.Code)
		}
	})
}

type mockUserStore struct {}

func (moch *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, nil
}

func (moch *mockUserStore) GetUserByID(id int) (*types.User, error) {
	return nil, nil
}

func (moch *mockUserStore) CreateUser(types.User, error) error {
	return nil
}
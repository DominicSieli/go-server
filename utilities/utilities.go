package utilities

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ParseJSON(request *http.Request, payload any) error {
	if request.Body == nil {
		return fmt.Errorf("missing request body")
	}

	return  json.NewDecoder(request.Body).Decode(payload)
}

func WriteJSON(response http.ResponseWriter, status int, v any) error {
	response.Header().Add("Content-Type", "application/json")
	response.WriteHeader(status)

	return json.NewEncoder(response).Encode(v)
}

func WriteError (response http.ResponseWriter, status int, err error) {
	WriteJSON(response, status, map[string]string{"error": err.Error()})
}
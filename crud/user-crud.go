package crud

import "net/http"

func LoginUser(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("Login User crud method called"))
}

func RegisterUser(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("Register User crud method called"))
}

func CreateUser(response http.ResponseWriter, request *http.Request) {}
func ReadUser(response http.ResponseWriter, request *http.Request)   {}
func UpdateUser(response http.ResponseWriter, request *http.Request) {}
func DeleteUser(response http.ResponseWriter, request *http.Request) {}

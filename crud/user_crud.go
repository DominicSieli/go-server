package crud

import "net/http"

func LoginUser(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("LoginUser crud method called"))
}

func CreateUser(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("CreateUser crud method called"))
}

func ReadUser(response http.ResponseWriter, request *http.Request)   {}
func UpdateUser(response http.ResponseWriter, request *http.Request) {}
func DeleteUser(response http.ResponseWriter, request *http.Request) {}

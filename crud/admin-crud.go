package crud

import "net/http"

func LoginAdmin(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("Login Admin crud method called"))
}

func RegisterAdmin(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("Register Admin crud method called"))
}

func CreateAdmin(response http.ResponseWriter, request *http.Request) {}
func ReadAdmin(response http.ResponseWriter, request *http.Request)   {}
func UpdateAdmin(response http.ResponseWriter, request *http.Request) {}
func DeleteAdmin(response http.ResponseWriter, request *http.Request) {}

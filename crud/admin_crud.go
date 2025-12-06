package crud

import "net/http"

func LoginAdmin(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("LoginAdmin crud method called"))
}

func CreateAdmin(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("CreateAdmin crud method called"))
}

func ReadAdmin(response http.ResponseWriter, request *http.Request)   {}
func UpdateAdmin(response http.ResponseWriter, request *http.Request) {}
func DeleteAdmin(response http.ResponseWriter, request *http.Request) {}

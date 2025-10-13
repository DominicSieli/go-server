package data

import "net/http"

func Login(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("Login crud method called"))
}

func Register(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("Register crud method called"))
}

func Create(response http.ResponseWriter, request *http.Request) {}
func Read(response http.ResponseWriter, request *http.Request)   {}
func Update(response http.ResponseWriter, request *http.Request) {}
func Delete(response http.ResponseWriter, request *http.Request) {}

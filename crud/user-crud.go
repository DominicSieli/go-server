package crud

import "net/http"

func UserLogin(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("UserLogin crud method called"))
}

func UserRegister(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("UserRegister crud method called"))
}

func UserCreate(response http.ResponseWriter, request *http.Request) {}
func UserRead(response http.ResponseWriter, request *http.Request)	 {}
func UserUpdate(response http.ResponseWriter, request *http.Request) {}
func UserDelete(response http.ResponseWriter, request *http.Request) {}

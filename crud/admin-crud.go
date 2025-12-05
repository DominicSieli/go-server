package crud

import "net/http"

func AdminLogin(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("AdminLogin crud method called"))
}

func AdminRegister(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("AdminRegister crud method called"))
}

func AdminCreate(response http.ResponseWriter, request *http.Request) {}
func AdminRead(response http.ResponseWriter, request *http.Request)   {}
func AdminUpdate(response http.ResponseWriter, request *http.Request) {}
func AdminDelete(response http.ResponseWriter, request *http.Request) {}

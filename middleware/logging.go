package middleware

import (
	"log"
	"net/http"
	"time"
)

type wrappedWriter struct {
	http.ResponseWriter
	statusCode int
}

func (response *wrappedWriter) WriteHeader(statusCode int) {
	response.ResponseWriter.WriteHeader(statusCode)
	response.statusCode = statusCode
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		start := time.Now()

		wrapped := &wrappedWriter{
			ResponseWriter: response,
			statusCode:		http.StatusOK,
		}

		next.ServeHTTP(wrapped, request)
		log.Println(wrapped.statusCode, request.Method, request.URL.Path, time.Since(start))
	})
}

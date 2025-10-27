package middleware

import "net/http"

type MiddleWare func(http.Handler) http.Handler

func CreateStack(args ...MiddleWare) MiddleWare {
	return func(next http.Handler) http.Handler {
		for i := len(args) - 1; i >= 0; i-- {
			middleware := args[i]
			next = middleware(next)
		}

		return next
	}
}

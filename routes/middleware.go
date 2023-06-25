package routes

import (
	"fmt"
	"net/http"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Middleware")
		fmt.Println(r.Method)
		fmt.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

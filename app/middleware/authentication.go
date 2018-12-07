package middleware

import (
	"net/http"	
	"go-app/app/response"
)

type Authentication struct {
}

func (amw Authentication) CheckAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-Session-Token")
		if token == "123" {
			next.ServeHTTP(w, r)
		} else {
			response.GlobalResponse{}.WithError(w, http.StatusForbidden, "Error", "Forbidden")
		}		
	})
}
package middleware

import (
	"go-app/app/models"
	"go-app/app/response"
	"net/http"
	"strings"

	"fmt"

	"github.com/jinzhu/gorm"
)

type Authentication struct {
	DB *gorm.DB
}

func (amw Authentication) CheckAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("X-Session-Token")
		if validToken := strings.Contains(authHeader, ";userId="); !validToken {
			response.GlobalResponse{}.WithError(w, http.StatusForbidden, "Error", "Forbidden")
			return
		}

		s := strings.Split(authHeader, ";userId=")
		token, userId := s[0], s[1]

		var user models.User
		// amw.db.Find(&user, "token = ? AND id = ?", token, userId)
		rows := amw.DB.Where(&user, "id = ?", "1").Value

		fmt.Println(rows)

		if token != "123" || userId != "444" {
			response.GlobalResponse{}.WithError(w, http.StatusForbidden, "Error", "Forbidden")
			return
		}
		next.ServeHTTP(w, r)
	})
}

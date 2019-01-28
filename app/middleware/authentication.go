package middleware

import (
	"go-app/app/models"
	"go-app/app/response"
	"net/http"
	"strings"

	"github.com/jinzhu/gorm"
)

type Authentication struct {
	DB   *gorm.DB
	User *models.User
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

		if err := amw.DB.Where("id = ? AND token = ?", userId, token).First(&amw.User).Error; err != nil {
			response.GlobalResponse{}.WithError(w, http.StatusForbidden, "Error", "Forbidden")
			return
		}

		next.ServeHTTP(w, r)
	})
}

package middleware

import (
	"go-app/app/helpers"
	"go-app/app/models"
	"go-app/app/response"
	"net/http"

	"github.com/jinzhu/gorm"
)

type Authentication struct {
	DB   *gorm.DB
	User *models.User
}

func (amw Authentication) CheckAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-Session-Token")

		if err := amw.DB.Where("token = ?", token).First(&amw.User).Error; err != nil || amw.User.Token == "" {
			helpers.ClearStruct(amw.User)
			response.GlobalResponse{}.WithError(w, http.StatusForbidden, "Error", "Forbidden")
			return
		}

		next.ServeHTTP(w, r)
	})
}

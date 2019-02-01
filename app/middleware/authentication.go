package middleware

import (
	// "go-app/app/helpers"
	"go-app/app/models"
	"go-app/app/response"
	"net/http"

	"github.com/jinzhu/gorm"
)

type Authentication struct {
	DB   *gorm.DB
}

func (amw Authentication) CheckAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-Session-Token")

		var userToken models.UserToken

		if err := amw.DB.Where("token = ?", token).First(&userToken).Error; err != nil || userToken.Active == 0 {
			response.GlobalResponse{}.WithError(w, http.StatusForbidden, "Error", "Forbidden")
			return
		}

		next.ServeHTTP(w, r)
	})
}

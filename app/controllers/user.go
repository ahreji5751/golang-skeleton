package controllers

import (
	"encoding/json"
	"go-app/app/helpers"
	"go-app/app/models"
	"net/http"

	validator "gopkg.in/validator.v2"
)

type UserController struct {
	super Controller
}

type UserRegisterSuccessResponse struct {
	ID    int64
	Email string
	Token string
}

func (uc UserController) Index(w http.ResponseWriter, req *http.Request) {
	// data := Data{"I'm from users index"}
	// uc.super.Response.WithJson(w, http.StatusOK, data)
}

func (uc UserController) Create(w http.ResponseWriter, req *http.Request) {
	var user models.User

	json.NewDecoder(req.Body).Decode(&user)

	if errs := validator.Validate(user); errs != nil {
		uc.super.Response.WithJson(w, http.StatusBadRequest, errs)
		return
	}

	user.EncryptedPassword, user.Token = helpers.GenerateToken(user.Password)

	if errInterface := uc.super.DB.Save(&user); errInterface.Error != nil {
		uc.super.Response.WithJson(w, http.StatusBadRequest, errInterface.Error)
		return
	}

	uc.super.Response.WithJson(w, http.StatusOK, UserRegisterSuccessResponse{user.ID, user.Email, user.Token})
}

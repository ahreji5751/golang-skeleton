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

type UserSuccessResponse struct {
	ID    int64
	Email string
	Token string
}

var user models.User

func (uc UserController) Index(w http.ResponseWriter, req *http.Request) {
	if err := uc.super.DB.Where("id = ?", uc.super.User.ID).First(&uc.super.User).Error; err != nil {
		uc.super.Response.WithJson(w, http.StatusOK, err)
	}

	uc.super.Response.WithJson(w, http.StatusOK, UserSuccessResponse{uc.super.User.ID, uc.super.User.Email, uc.super.User.Token})
}

func (uc UserController) Create(w http.ResponseWriter, req *http.Request) {
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

	uc.super.Response.WithJson(w, http.StatusOK, UserSuccessResponse{user.ID, user.Email, user.Token})
}

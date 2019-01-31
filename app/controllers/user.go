package controllers

import (
	"encoding/json"
	"go-app/app/helpers"
	"go-app/app/models"
	"net/http"

	validator "gopkg.in/validator.v2"
)

var (
	creationValidator  = validator.NewValidator()
	signingInValidator = validator.NewValidator()
)

type UserController struct {
	super Controller
}

type UserSuccessResponse struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
	Token string `json:"token"`
}

type EmptyResponse struct {
}

func (uc UserController) Index(w http.ResponseWriter, req *http.Request) {
	if err := uc.super.DB.Where("id = ?", uc.super.User.ID).First(&uc.super.User).Error; err != nil {
		uc.super.Response.WithJson(w, http.StatusOK, err)
	}

	uc.super.Response.WithJson(w, http.StatusOK, UserSuccessResponse{uc.super.User.ID, uc.super.User.Email, uc.super.User.Name, uc.super.User.Token})
}

func (uc UserController) Create(w http.ResponseWriter, req *http.Request) {
	var user models.User

	json.NewDecoder(req.Body).Decode(&user)

	creationValidator.SetTag("create")

	if errs := creationValidator.Validate(user); errs != nil {
		uc.super.Response.WithJson(w, http.StatusBadRequest, errs)
		return
	}

	if err := uc.super.DB.Where("email = ?", user.Email).First(&user).Error; err == nil {
		uc.super.Response.WithError(w, http.StatusBadRequest, "Error", "User with this email already exist!")
		return
	}

	user.EncryptedPassword, user.Token = helpers.GeneratePassword(user.Password), helpers.GenerateToken()

	if err := uc.super.DB.Create(&user).Error; err != nil {
		uc.super.Response.WithJson(w, http.StatusBadRequest, err)
		return
	}

	uc.super.Response.WithJson(w, http.StatusOK, UserSuccessResponse{user.ID, user.Email, user.Name, user.Token})
}

func (uc UserController) Login(w http.ResponseWriter, req *http.Request) {
	var user models.User

	json.NewDecoder(req.Body).Decode(&user)

	signingInValidator.SetTag("signin")

	if errs := signingInValidator.Validate(user); errs != nil {
		uc.super.Response.WithJson(w, http.StatusBadRequest, errs)
		return
	}

	requestPassword := user.Password

	if err := uc.super.DB.Where("email = ?", user.Email).First(&user).Error; err != nil {
		uc.super.Response.WithError(w, http.StatusBadRequest, "Error", "User does not exist!")
		return
	}

	if !helpers.PasswordIsValid(requestPassword, user.EncryptedPassword) {
		uc.super.Response.WithError(w, http.StatusBadRequest, "Error", "Incorrect password!")
		return
	}

	user.Token = helpers.GenerateToken()

	if err := uc.super.DB.Save(&user).Error; err != nil {
		uc.super.Response.WithJson(w, http.StatusBadRequest, err)
		return
	}

	uc.super.Response.WithJson(w, http.StatusOK, UserSuccessResponse{user.ID, user.Email, user.Name, user.Token})
}

func (uc UserController) Logout(w http.ResponseWriter, req *http.Request) {
	if err := uc.super.DB.Where("id = ?", uc.super.User.ID).First(&uc.super.User).Error; err != nil {
		uc.super.Response.WithJson(w, http.StatusBadRequest, err)
	}

	uc.super.User.Token = ""

	if err := uc.super.DB.Save(&uc.super.User).Error; err != nil {
		uc.super.Response.WithJson(w, http.StatusBadRequest, err)
		return
	}

	uc.super.Response.WithJson(w, http.StatusOK, EmptyResponse{})
}

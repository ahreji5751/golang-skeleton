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

type UserResource struct {
	ID    uint64 `deepcopier:"field:ID" json:"id"`
	Email string `deepcopier:"field:Email" json:"email"`
	Name  string `deepcopier:"field:Name" json:"name"`
	Token string `deepcopier:"field:Token" json:"token"`
}

func (uc UserController) Index(w http.ResponseWriter, req *http.Request) {
	uc.super.Response.WithJson(w, http.StatusOK, &UserResource{uc.super.User.ID, uc.super.User.Email, uc.super.User.Name, uc.super.Token})
}

func (uc UserController) Create(w http.ResponseWriter, req *http.Request) {
	var user models.User
	var userToken models.UserToken

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

	user.EncryptedPassword, userToken = helpers.GeneratePassword(user.Password), models.UserToken{Token: helpers.GenerateToken()}

	if err := uc.super.DB.Create(&user).Error; err != nil {
		uc.super.Response.WithJson(w, http.StatusBadRequest, err)
		return
	}

	if err := uc.super.DB.Model(&user).Association("Tokens").Append(userToken).Error; err != nil {
		uc.super.Response.WithJson(w, http.StatusBadRequest, err)
		return
	}

	uc.super.Response.WithJson(w, http.StatusOK, &UserResource{user.ID, user.Email, user.Name, userToken.Token})
}

func (uc UserController) Login(w http.ResponseWriter, req *http.Request) {
	var user models.User
	var userToken models.UserToken

	json.NewDecoder(req.Body).Decode(&user)

	signingInValidator.SetTag("signin")

	if errs := signingInValidator.Validate(user); errs != nil {
		uc.super.Response.WithJson(w, http.StatusBadRequest, errs)
		return
	}

	if err := uc.super.DB.Where("email = ?", user.Email).First(&user).Error; err != nil {
		uc.super.Response.WithError(w, http.StatusBadRequest, "Error", "User does not exist!")
		return
	}

	if !helpers.PasswordIsValid(user.Password, user.EncryptedPassword) {
		uc.super.Response.WithError(w, http.StatusBadRequest, "Error", "Incorrect password!")
		return
	}

	if err := uc.super.DB.Where("user_id = ? AND active = 0", user.ID).First(&userToken).Error; err == nil {
		userToken.Active = 1
		if err := uc.super.DB.Save(&userToken).Error; err != nil {
			uc.super.Response.WithJson(w, http.StatusBadRequest, err)
			return
		}
		uc.super.Response.WithJson(w, http.StatusOK, &UserResource{user.ID, user.Email, user.Name, userToken.Token})
		return
	}

	userToken = models.UserToken{Token: helpers.GenerateToken()}

	if err := uc.super.DB.Model(&user).Association("Tokens").Append(userToken).Error; err != nil {
		uc.super.Response.WithJson(w, http.StatusBadRequest, err)
		return
	}

	uc.super.Response.WithJson(w, http.StatusOK, &UserResource{user.ID, user.Email, user.Name, userToken.Token})
}

func (uc UserController) Logout(w http.ResponseWriter, req *http.Request) {
	var userToken models.UserToken

	if err := uc.super.DB.Where("token = ?", uc.super.Token).First(&userToken).Error; err != nil {
		uc.super.Response.WithJson(w, http.StatusBadRequest, err)
	}

	userToken.Active = 0

	if err := uc.super.DB.Save(&userToken).Error; err != nil {
		uc.super.Response.WithJson(w, http.StatusBadRequest, err)
		return
	}

	uc.super.Response.WithJson(w, http.StatusOK, &UserResource{})
}

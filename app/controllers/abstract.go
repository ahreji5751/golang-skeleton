package controllers

import (
	"go-app/app/models"
	"go-app/app/response"
	"net/http"
	"reflect"

	"github.com/jinzhu/gorm"
)

type Controller struct {
	Name           string
	Action         string
	ResponseWriter http.ResponseWriter
	Request        *http.Request
	Response       *response.GlobalResponse
	DB             *gorm.DB
	User           *models.User
	Token          string
}

func (c Controller) Run(controllerType interface{}) {
	reflect.ValueOf(controllerType).MethodByName(c.Action).Call([]reflect.Value{reflect.ValueOf(c.ResponseWriter), reflect.ValueOf(c.Request)})
}

func (c Controller) Migrate(models []interface{}) {
	for _, model := range models {
		if !c.DB.HasTable(model) {
			c.DB.AutoMigrate(model)
		}
	}
}

func Call(controllerName string, actionName string, w http.ResponseWriter, req *http.Request, db *gorm.DB, withUser bool) {
	user, token := getCurrentUser(req, db, withUser)
	c := Controller{controllerName, actionName, w, req, &response.GlobalResponse{}, db, user, token}
	mapper := map[string][]interface{}{
		"User": []interface{}{UserController{c}, models.User{}, models.UserToken{}},
	}
	c.Migrate([]interface{}{mapper[controllerName][1], mapper[controllerName][2]})
	c.Run(mapper[controllerName][0])
}

func getCurrentUser(req *http.Request, db *gorm.DB, withUser bool) (*models.User, string) {
	if !withUser {
		return nil, ""
	}

	var user models.User
	var userToken models.UserToken

	token := req.Header.Get("X-Session-Token")

	if err := db.Where("token = ?", token).First(&userToken).Error; err != nil {
		panic("User for this account not found!")
	}

	if err := db.Model(&userToken).Related(&user).Error; err != nil {
		panic("User for this account not found!")
	}

	return &user, userToken.Token
}

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
}

func (c Controller) Run(controllerType interface{}) {
	reflect.ValueOf(controllerType).MethodByName(c.Action).Call([]reflect.Value{reflect.ValueOf(c.ResponseWriter), reflect.ValueOf(c.Request)})
}

func (c Controller) Migrate(model interface{}) {
	if !c.DB.HasTable(model) {
		c.DB.AutoMigrate(model)
	}
}

func Call(controllerName string, actionName string, w http.ResponseWriter, req *http.Request, db *gorm.DB, user *models.User) {
	c := Controller{controllerName, actionName, w, req, &response.GlobalResponse{}, db, user}
	mapper := map[string][]interface{}{
		"User": []interface{}{UserController{c}, models.User{}},
	}
	c.Migrate(mapper[controllerName][1])
	c.Run(mapper[controllerName][0])
}

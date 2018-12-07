package controllers

import (
	"reflect"
	"net/http"
	"go-app/app/response"
	"go-app/app/requests"
)

type Controller struct {
	Name           string   
	Action         string 
	ResponseWriter http.ResponseWriter 
	Request        *http.Request   
	Response	   *response.GlobalResponse
}

func (c Controller) Run(controllerType interface{}, requestType interface{}) {
	if isValid := c.requestIsValid(requestType); isValid {
		reflect.ValueOf(controllerType).MethodByName(c.Action).Call([]reflect.Value{reflect.ValueOf(c.ResponseWriter), reflect.ValueOf(c.Request)})
	}
}

func (c Controller) requestIsValid(requestType interface{}) bool {
	if _, validatorExist := reflect.TypeOf(requestType).MethodByName(c.Action); validatorExist {
		return reflect.ValueOf(requestType).MethodByName(c.Action).Call([]reflect.Value{})[0].Bool()
	} 
	return true
}

func Call(controllerName string, actionName string, w http.ResponseWriter, req *http.Request) {
	c := Controller{controllerName, actionName, w, req, &response.GlobalResponse{}}
	mapper := map[string][]interface{} {
		"Product": []interface{}{Product{c}, requests.Product{w, req, &response.GlobalResponse{}}},
		"User": []interface{}{User{c}, requests.User{w, req, &response.GlobalResponse{}}},
	}
	c.Run(mapper[controllerName][0], mapper[controllerName][1])
}
package requests

import (
	// "github.com/thedevsaddam/govalidator"
	"go-app/app/response"
	"net/http"
)

type User struct {
	ResponseWriter  http.ResponseWriter
	Request  		*http.Request
	Response        *response.GlobalResponse
}



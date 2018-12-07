package requests

import (
	"github.com/thedevsaddam/govalidator"
	"go-app/app/response"
	"net/http"
)

type Product struct {
	ResponseWriter  http.ResponseWriter
	Request  		*http.Request
	Response        *response.GlobalResponse
}

func (pr Product) Create() bool {
	return Request{response: pr.Response, w: pr.ResponseWriter}.BuildValidatorAndValidate(
		pr.Request, 
		govalidator.MapData{
			"username": []string{"required", "between:3,8"},
		}, 
		govalidator.MapData{
			"username": []string{"required:Username is required", "between:Must be 3 to 8 symbols"},
		},
	)
}


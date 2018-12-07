package requests

import (
	"github.com/thedevsaddam/govalidator"
	"reflect"
	"go-app/app/response"
	"net/http"
)

type Request struct {
	opts     govalidator.Options
	response *response.GlobalResponse
	w        http.ResponseWriter
}

func (r Request) BuildValidatorAndValidate(req *http.Request, rules govalidator.MapData, messages govalidator.MapData) bool {
	return Request{
		govalidator.Options{
			Request:         req,       
			Rules:           rules,    
			Messages:        messages,
			RequiredDefault: true,    	
		},
		r.response, 
		r.w,	
	}.validate()
}

func (r Request) validate() bool {
	v := govalidator.New(r.opts)
	e := v.Validate()
	a := map[string]interface{}{"validationError": e}	
	if !isEmpty(e)  {
		r.response.WithJson(r.w, http.StatusBadRequest, a)
	}
	return isEmpty(e)
}



func isEmpty(x interface{}) bool {
	rt := reflect.TypeOf(x)
	if rt == nil {
		return true
	}
	rv := reflect.ValueOf(x)
	switch rv.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice:
		return rv.Len() == 0
	}
	return reflect.DeepEqual(x, reflect.Zero(rt).Interface())
}
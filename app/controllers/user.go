package controllers

import (
	"net/http"
)

type User struct {
	super Controller
}

func (c User) Index(w http.ResponseWriter, req *http.Request)  {
	data := Data{"I'm from users index"}
	c.super.Response.WithJson(w, http.StatusOK, data)
}
package controllers

import (
	"net/http"
)

type Product struct {
	super Controller
}

type Data struct {
	Test string
}

func (c Product) Index(w http.ResponseWriter, req *http.Request)  {
	data := Data{"I'm from products index"}
	c.super.Response.WithJson(w, http.StatusOK, data)
	// c.Ac.Response.WithError(w, 403, "Error", "Unauthorized")
}

func (c Product) Create(w http.ResponseWriter, req *http.Request)  {
	data := Data{"I'm from products post"}
	c.super.Response.WithJson(w, http.StatusOK, data)
	// c.Ac.Response.WithError(w, 403, "Unauthorized")
}
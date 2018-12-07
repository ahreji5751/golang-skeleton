package routes

import (
	"go-app/app/controllers"
	"go-app/app/middleware"
	"github.com/gorilla/mux"
	"net/http"
	"path/filepath"
	"os"
	"log"
)

type Router struct {
	Mr 	*mux.Router
}

func (r Router) Route() {
	pr := r.Mr.PathPrefix("/products").Subrouter()

	pr.HandleFunc("", func (w http.ResponseWriter, req *http.Request) { controllers.Call("Product", "Index", w, req) }).Methods("GET")
	pr.HandleFunc("", func (w http.ResponseWriter, req *http.Request) { controllers.Call("Product", "Create", w, req) }).Methods("POST")

	ur := r.Mr.PathPrefix("/users").Subrouter()

	ur.Use(middleware.Authentication{}.CheckAuth)
	ur.HandleFunc("", func (w http.ResponseWriter, req *http.Request) { controllers.Call("User", "Index", w, req) }).Methods("GET")

	r.Mr.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(currentFilePath() + "/public"))))
	r.Mr.NotFoundHandler = http.HandlerFunc(notFound)
}

func notFound(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, currentFilePath() + "/public/404.html")	
}

func currentFilePath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
    if err != nil {
        log.Fatal(err)
	}
	return dir
}
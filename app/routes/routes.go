package routes

import (
	"go-app/app/controllers"
	"go-app/app/middleware"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type Router struct {
	Mr *mux.Router
}

func (r Router) Route(db *gorm.DB) {
	ur := r.Mr.PathPrefix("/users").Subrouter()

	ur.Use(middleware.Authentication{db}.CheckAuth)
	ur.HandleFunc("", func(w http.ResponseWriter, req *http.Request) { controllers.Call("User", "Index", w, req, db) }).Methods("GET")
	ur.HandleFunc("/register", func(w http.ResponseWriter, req *http.Request) { controllers.Call("User", "Create", w, req, db) }).Methods("POST")

	r.Mr.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(currentFilePath()+"/public"))))
	r.Mr.NotFoundHandler = http.HandlerFunc(notFound)
}

func notFound(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, currentFilePath()+"/public/404.html")
}

func currentFilePath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

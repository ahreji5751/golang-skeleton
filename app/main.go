package main

import (
	"fmt"
	"go-app/app/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	router := mux.NewRouter()
	appRouter := routes.Router{router}

	db, err := gorm.Open("mysql", "root:arimib@/go_test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	appRouter.Route(db)

	fmt.Println("Listening on port: 8080")

	log.Fatal(http.ListenAndServe(":8080", router))
}

package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"fmt"
	"go-app/app/routes"
)

func main() {
	router := mux.NewRouter()
	appRouter := routes.Router{router}
	appRouter.Route()

	fmt.Println("Listening on port: 8080")

	log.Fatal(http.ListenAndServe(":8080", router))
}
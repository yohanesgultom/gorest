package main

import (
	"github.com/gorilla/mux"
	"os"
	"fmt"
	"net/http"
	"main/controllers"
	"main/app"
)

func main() {
	
	router := mux.NewRouter()
	auth := app.JwtAuthentication
	router.HandleFunc("/", controllers.Hello).Methods("GET")
	router.HandleFunc("/users", controllers.Register).Methods("POST")	
	router.HandleFunc("/auth/login", controllers.Login).Methods("POST")	
	router.Handle("/secured", auth(http.HandlerFunc(controllers.Secured))).Methods("GET")

	port := os.Getenv("app_port")
	if port == "" {
		port = "8000"
	}
	
	fmt.Println("Listening on " + port + "..")
	e := http.ListenAndServe(":" + port, router)
	if e != nil {
		fmt.Println(e)
	}

}
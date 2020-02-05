package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"go-login/controller"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/register", controller.RegisterHandler).
		Methods("POST")
	r.HandleFunc("/login", controller.LoginHandler).
		Methods("POST")
	r.HandleFunc("/profile", controller.ProfileHandler).
		Methods("GET")
	fmt.Println("server connected")
	log.Fatal(http.ListenAndServe(":8080", r))

}

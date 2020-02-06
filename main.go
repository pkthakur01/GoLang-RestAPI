package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"go-login/controller"
	logger "go-login/utils"
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
	logger.GeneralLogger.Println("Server Connected and Running at PORT 8081")
	log.Fatal(http.ListenAndServe(":8080", r))

}

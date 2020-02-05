package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title       string `json:"Title"`
	Description string `json:"desc"`
	Content     string `json:"content"`
}

type Atricles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	article := Atricles{
		Article{Title: "Test Title", Description: "Nothing to say", Content: "Hello Prakash"},
	}
	fmt.Println("All aticles will be printed")
	json.NewEncoder(w).Encode(article)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hey I'm testing my first Go app")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}

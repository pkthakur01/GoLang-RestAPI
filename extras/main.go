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

// mymap_1 := map[int]string{
// 	12: "Prakash",
// 	13: "WhoCares",
// }

type Atricles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {

	map_2 := map[string]string{

		"90": "Dog",
		"91": "Cat",
		"92": "Cow",
		"93": "Bird",
		"94": "Rabbit",
	}
	fmt.Println(map_2)
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

func sayHello() {
	fmt.Println("Hello Prakash Bhai")
}

func main() {
	sayHello()
	handleRequests()

}

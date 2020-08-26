package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Test title", Desc: "Demo Desc", Content: "Content is here"},
	}
	fmt.Fprintf(w, "Endpoint: allArticles")
	json.NewEncoder(w).Encode(articles)
	fmt.Println("Endpoint Hit: allArticles")
}

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Homepage Endpoint")
	fmt.Println("Endpoint: Homepage")
}

func handleRequest() {
	http.HandleFunc("/", homepage)
	http.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":5000", nil))
}

func main() {
	handleRequest()
}

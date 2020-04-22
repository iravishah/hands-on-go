package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article

func main() {
	fmt.Println("Hands on Rest API")
	loadArticals()
	handlerRequests()
}

func loadArticals() {
	Articles = []Article{
		Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
}

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func getAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func handlerRequests() {
	apiRouter := mux.NewRouter().StrictSlash(true)
	apiRouter.HandleFunc("/", homepage)
	apiRouter.HandleFunc("/articles", getAllArticles)
	log.Fatal(http.ListenAndServe(":9999", apiRouter))
}

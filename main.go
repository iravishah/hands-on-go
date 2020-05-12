package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Id      string
	Title   string
	Desc    string
	Content string
}

type Addition struct {
	Sum int
}

var Articles []Article

func main() {
	fmt.Println("Hands on Rest API")
	loadArticals()
	handlerRequests()
}

func loadArticals() {
	Articles = []Article{
		Article{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
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

func getArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func createArticle(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var article Article
	json.Unmarshal(body, &article)
	Articles = append(Articles, article)
	json.NewEncoder(w).Encode(article)
}

func removeArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for index, article := range Articles {
		if article.Id == key {
			Articles = append(Articles[:index], Articles[index+1:]...)
		}
	}
}

func updateArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var article1 Article

	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &article1)

	for index, article := range Articles {
		if article.Id == key {
			Articles[index] = article1
		}
	}
	json.NewEncoder(w).Encode(Articles)
}

func addition(w http.ResponseWriter, r *http.Request) {
	add := Addition{Sum: 2 + 2}
	json.NewEncoder(w).Encode(add)
}

func insert(w http.ResponseWriter, r *http.Request) {
	connect.connectMongo()
}

func handlerRequests() {
	apiRouter := mux.NewRouter().StrictSlash(true)
	apiRouter.HandleFunc("/", homepage)
	apiRouter.HandleFunc("/articles", getAllArticles).Methods("GET")
	apiRouter.HandleFunc("/articles/{id}", getArticle).Methods("GET")
	apiRouter.HandleFunc("/article", createArticle).Methods("POST")
	apiRouter.HandleFunc("/article/{id}", removeArticle).Methods("DELETE")
	apiRouter.HandleFunc("/article/{id}", updateArticle).Methods("PUT")
	apiRouter.HandleFunc("/add", addition).Methods("GET")
	apiRouter.HandleFunc("/insert", insert).Methods("POST")

	log.Fatal(http.ListenAndServe(":9999", apiRouter))
}

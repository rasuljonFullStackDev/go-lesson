package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Article struct {
	Title   string "json:title"
	Body    string "json:body"
	Comment string "json:comment"
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Hello", Body: "Hello world", Comment: "Comment"},
		Article{Title: "Hello 2", Body: "Hello world 2", Comment: "Comment 2"},
		Article{Title: "Hello 3", Body: "Hello world 3", Comment: "Comment 3"},
	}
	fmt.Println("Endpoint Hit: allArticles")
	json.NewEncoder(w).Encode(articles)
	// articles := Articles{
	// 	Article{Title: "Hello", Body: "Hello world", Comment: "Comment"},
	// 	Article{Title: "Hello 2", Body: "Hello world 2", Comment: "Comment 2"},
	// }
	// fmt.Println("Endpoint Hit: allArticles")
	// json.NewDecoder(w).Encode(articles)
	// articles := Articles{
	// 	Article{Title: "Hello", Body: "Hello world", Comment: "Comment"},
	// 	Article{Title: "Hello 2", Body: "Hello world 2", Comment: "Comment 2"},
	// }
	// fmt.Println("Endpoint Hit: allArticles")
	// json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}
func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

// doim main bolishi kerak
func main() {
	handleRequests()
}

// qaysi papka bolsa osha yoziladi asosiyga main yoziladi
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
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
func notFound(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Not Found")
}
func storeArticle(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Article{Title: "Hello", Body: "Hello world", Comment: "Comment"})
}
func handleRequests() {
	// 1 dars
	// http.HandleFunc("/", homePage)
	// http.HandleFunc("/articles", allArticles)
	// log.Fatal(http.ListenAndServe(":8080", nil))
	// 2 dars
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	myRouter.HandleFunc("/articles", storeArticle).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
	// not foudni beradi
	myRouter.NotFoundHandler = http.HandlerFunc(notFound)

}

// doim main bolishi kerak
func main() {
	handleRequests()
}

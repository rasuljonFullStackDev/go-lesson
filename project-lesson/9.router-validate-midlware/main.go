package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// | Tushuncha                   | Tavsif              | Misol                                     |
// | --------------------------- | ------------------- | ----------------------------------------- |
// | `mux.NewRouter()`           | Router yaratadi     | `r := mux.NewRouter()`                    |
// | `HandleFunc(path, handler)` | Yo‘lni belgilaydi   | `r.HandleFunc("/users", getUsers)`        |
// | `Methods("GET")`            | HTTP method         | `.Methods("POST")`                        |
// | `mux.Vars(r)`               | URL parametrlari    | `id := mux.Vars(r)["id"]`                 |
// | `r.Use()`                   | Middleware qo‘shish | `r.Use(loggingMiddleware)`                |
// | `PathPrefix("/api")`        | Subrouter           | `api := r.PathPrefix("/api").Subrouter()` |

// go get -u github.com/gorilla/mux
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("So‘rov:", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func main() {
	fmt.Println("loyiha ishga tushdi")
	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()
	r.Use(loggingMiddleware)

	api.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Bosh sahifa")
	}).Methods("GET")

	api.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Biz haqimizda")
	}).Methods("GET")

	api.HandleFunc("/user/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		fmt.Fprintf(w, "Foydalanuvchi ID: %s", id)
	}).Methods("GET")

	api.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query().Get("q")
		fmt.Fprintf(w, "Qidiruv: %s", q)
	}).Methods("GET")
	api.HandleFunc("/product", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Faqat POST", http.StatusMethodNotAllowed)
			return
		}

		type Product struct {
			Name  string `json:"name"`
			Price int    `json:"price"`
		}
		var p Product
		json.NewDecoder(r.Body).Decode(&p)

		fmt.Fprintf(w, "Qabul qilindi: %s – %d so‘m", p.Name, p.Price)
	}).Methods("POST")
	// 🔥 404 Handler:
	api.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]any{
			"error":   "not_found",
			"message": "Bu sahifa topilmadi",
			"path":    r.URL.Path,
		})
	})
	fmt.Println("Server: http://localhost:8080")
	http.ListenAndServe(":8080", r)

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintln(w, "Bosh sahifa")
	// })
	// http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintln(w, "Biz haqimizda")
	// })

	// fmt.Println("Server: http://localhost:8080")
	// http.ListenAndServe(":8080", nil)
}

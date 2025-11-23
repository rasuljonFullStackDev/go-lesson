package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

var (
	db  *sql.DB
	tpl *template.Template
)

func main() {
	// DB connection
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_PORT", "5432"),
		getEnv("DB_USER", "postgres"),
		getEnv("DB_PASSWORD", "12345678"),
		getEnv("DB_NAME", "shopdb"),
	)
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// Templates with FuncMap
	tpl = template.Must(template.New("").Funcs(template.FuncMap{
		"dollar": func(cents int) string {
			return fmt.Sprintf("%.2f", float64(cents)/100)
		},
	}).ParseGlob("templates/*.html"))

	// Tables
	createTables()

	// Routes
	// User CRUD
	http.HandleFunc("/", userIndex)
	http.HandleFunc("/user/new", userNew)
	http.HandleFunc("/user/create", userCreate)
	http.HandleFunc("/user/edit", userEdit)
	http.HandleFunc("/user/update", userUpdate)
	http.HandleFunc("/user/delete", userDelete)

	// Product CRUD
	http.HandleFunc("/products", productIndex)
	http.HandleFunc("/product/new", productNew)
	http.HandleFunc("/product/create", productCreate)
	http.HandleFunc("/product/edit", productEdit)
	http.HandleFunc("/product/update", productUpdate)
	http.HandleFunc("/product/delete", productDelete)

	log.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

// ===== HELPERS =====
func getEnv(key, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	return val
}

// ===== DB =====
func createTables() {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username TEXT UNIQUE NOT NULL,
		email TEXT NOT NULL
	);
	CREATE TABLE IF NOT EXISTS products (
		id SERIAL PRIMARY KEY,
		name TEXT,
		price INT
	);
	`)
	if err != nil {
		log.Fatal(err)
	}
}

// ===== USER CRUD =====
func userIndex(w http.ResponseWriter, r *http.Request) {
	rows, _ := db.Query("SELECT id, username, email FROM users ORDER BY id DESC")
	defer rows.Close()
	var users []map[string]interface{}
	for rows.Next() {
		var id int
		var username, email string
		rows.Scan(&id, &username, &email)
		users = append(users, map[string]interface{}{"ID": id, "Username": username, "Email": email})
	}
	tpl.ExecuteTemplate(w, "user_index.html", users)
}

func userNew(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "user_new.html", nil)
}

func userCreate(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	email := r.FormValue("email")
	db.Exec("INSERT INTO users (username,email) VALUES ($1,$2)", username, email)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func userEdit(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	var username, email string
	db.QueryRow("SELECT username,email FROM users WHERE id=$1", id).Scan(&username, &email)
	data := map[string]interface{}{"ID": id, "Username": username, "Email": email}
	tpl.ExecuteTemplate(w, "user_edit.html", data)
}

func userUpdate(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	username := r.FormValue("username")
	email := r.FormValue("email")
	db.Exec("UPDATE users SET username=$1,email=$2 WHERE id=$3", username, email, id)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func userDelete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	db.Exec("DELETE FROM users WHERE id=$1", id)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// ===== PRODUCT CRUD =====
func productIndex(w http.ResponseWriter, r *http.Request) {
	rows, _ := db.Query("SELECT id, name, price FROM products ORDER BY id DESC")
	defer rows.Close()
	var products []map[string]interface{}
	for rows.Next() {
		var id, price int
		var name string
		rows.Scan(&id, &name, &price)
		products = append(products, map[string]interface{}{"ID": id, "Name": name, "Price": price})
	}
	tpl.ExecuteTemplate(w, "product_index.html", products)
}

func productNew(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "product_new.html", nil)
}

func productCreate(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	price, _ := strconv.Atoi(r.FormValue("price"))
	db.Exec("INSERT INTO products (name,price) VALUES ($1,$2)", name, price)
	http.Redirect(w, r, "/products", http.StatusSeeOther)
}

func productEdit(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	var name string
	var price int
	db.QueryRow("SELECT name,price FROM products WHERE id=$1", id).Scan(&name, &price)
	data := map[string]interface{}{"ID": id, "Name": name, "Price": price}
	tpl.ExecuteTemplate(w, "product_edit.html", data)
}

func productUpdate(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	name := r.FormValue("name")
	price, _ := strconv.Atoi(r.FormValue("price"))
	db.Exec("UPDATE products SET name=$1,price=$2 WHERE id=$3", name, price, id)
	http.Redirect(w, r, "/products", http.StatusSeeOther)
}

func productDelete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	db.Exec("DELETE FROM products WHERE id=$1", id)
	http.Redirect(w, r, "/products", http.StatusSeeOther)
}

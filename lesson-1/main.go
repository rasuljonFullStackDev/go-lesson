package main

import (
    "database/sql"
    "encoding/json"
    "log"
    "net/http"
    "os"

    _ "modernc.org/sqlite" // ✅ CGO kerak emas
    "github.com/gorilla/mux"
)

type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

var db *sql.DB

func main() {
    // 1️⃣ Baza faylini tekshirish yoki yaratish
    dbFile := "./users.db"
    if _, err := os.Stat(dbFile); os.IsNotExist(err) {
        log.Println("❗ Database topilmadi, yangisini yaratamiz...")
        file, err := os.Create(dbFile)
        if err != nil {
            log.Fatal(err)
        }
        file.Close()
    }

    // 2️⃣ Baza bilan ulan
    var err error
    db, err = sql.Open("sqlite", dbFile) // ✅ "sqlite3" emas, "sqlite"
    if err != nil {
        log.Fatal(err)
    }

    // 3️⃣ Jadval yaratish
    createTable()

    // 4️⃣ Router
    r := mux.NewRouter()
    r.HandleFunc("/users", getUsers).Methods("GET")
    r.HandleFunc("/users/{id}", getUser).Methods("GET")
    r.HandleFunc("/users", createUser).Methods("POST")
    r.HandleFunc("/users/{id}", updateUser).Methods("PUT")
    r.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")

    log.Println("🚀 Server ishga tushdi: http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}

func createTable() {
    query := `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        email TEXT NOT NULL UNIQUE
    );
    `
    _, err := db.Exec(query)
    if err != nil {
        log.Fatal("Jadval yaratishda xato:", err)
    }
}

func getUsers(w http.ResponseWriter, r *http.Request) {
    rows, err := db.Query("SELECT id, name, email FROM users")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var users []User
    for rows.Next() {
        var u User
        rows.Scan(&u.ID, &u.Name, &u.Email)
        users = append(users, u)
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}

func getUser(w http.ResponseWriter, r *http.Request) {
    id := mux.Vars(r)["id"]
    var u User
    err := db.QueryRow("SELECT id, name, email FROM users WHERE id = ?", id).Scan(&u.ID, &u.Name, &u.Email)
    if err != nil {
        http.Error(w, "Foydalanuvchi topilmadi", http.StatusNotFound)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(u)
}

func createUser(w http.ResponseWriter, r *http.Request) {
    var u User
    json.NewDecoder(r.Body).Decode(&u)

    result, err := db.Exec("INSERT INTO users (name, email) VALUES (?, ?)", u.Name, u.Email)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    id, _ := result.LastInsertId()
    u.ID = int(id)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(u)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
    id := mux.Vars(r)["id"]
    var u User
    json.NewDecoder(r.Body).Decode(&u)

    _, err := db.Exec("UPDATE users SET name = ?, email = ? WHERE id = ?", u.Name, u.Email, id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(u)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
    id := mux.Vars(r)["id"]
    _, err := db.Exec("DELETE FROM users WHERE id = ?", id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusNoContent)
}

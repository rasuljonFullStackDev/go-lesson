package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct {
	Name  string
	Price int
	Color string
}

func withAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token != "Bearer 12345" {
			http.Error(w, "Ruxsat yo‘q!", http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}

func main() {
	fmt.Println("loyiha ishga tushdi")
	data := [...]Product{
		{Name: "apple", Price: 100, Color: "red"},
		{Name: "apple", Price: 100, Color: "red"},
		{Name: "apple", Price: 100, Color: "red"},
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})
	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(data)
	})
	http.HandleFunc("/secret", withAuth(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Xush kelibsiz, maxfiy sahifa!")
	}))
	http.HandleFunc("/param", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		fmt.Fprint(w, name)
		//  r.Header
		// r.Method
	})
	http.HandleFunc("/json", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Hello, Go!",
			"status":  "success",
		})
	})
	// fmt.Println("URL:", r.URL.Path)
	// fmt.Println("Host:", r.Host)
	// fmt.Println("Method:", r.Method)
	// fmt.Println("RemoteAddr:", r.RemoteAddr)
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Faqat POST so‘rovga ruxsat", http.StatusMethodNotAllowed)
			return
		}

		username := r.FormValue("username")
		password := r.FormValue("password")

		if username != "admin" || password != "1234" {
			http.Error(w, "Login yoki parol xato", http.StatusUnauthorized)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Tizimga muvaffaqiyatli kirdingiz ✅")
	})
	http.ListenAndServe(":8080", nil)
}

// | Amal                      | Funksiya                | Tavsif                        |
// | ------------------------- | ----------------------- | ----------------------------- |
// | `r.Method`                | So‘rov turi (GET, POST) | So‘rovni aniqlaydi            |
// | `r.URL.Query()`           | Query param             | `?key=value`                  |
// | `r.FormValue("x")`        | Form ma’lumoti          | POST form                     |
// | `r.Header.Get("x")`       | Header olish            | `Authorization`, `User-Agent` |
// | `w.Header().Set()`        | Javob header            | `Content-Type`                |
// | `json.NewDecoder(r.Body)` | JSON o‘qish             | API request body              |
// | `json.NewEncoder(w)`      | JSON yuborish           | API javobi                    |
// | `r.ParseMultipartForm()`  | Fayl yuklash            | Form-data                     |
// | `http.Error()`            | Status + xabar yuborish | `404`, `401` va h.k.          |

// | Kod                              | Ma’nosi | Tavsif             |
// | -------------------------------- | ------- | ------------------ |
// | `http.StatusOK`                  | 200     | Hammasi joyida     |
// | `http.StatusCreated`             | 201     | Resurs yaratildi   |
// | `http.StatusBadRequest`          | 400     | Noto‘g‘ri so‘rov   |
// | `http.StatusUnauthorized`        | 401     | Ruxsat yo‘q        |
// | `http.StatusForbidden`           | 403     | Ta’qiqlangan       |
// | `http.StatusNotFound`            | 404     | Topilmadi          |
// | `http.StatusConflict`            | 409     | Mojaro (duplicate) |
// | `http.StatusInternalServerError` | 500     | Ichki xato         |

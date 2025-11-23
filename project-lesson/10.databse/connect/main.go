package main

// go get -u gorm.io/gorm
// go get -u gorm.io/driver/postgres
import (
	"context"
	"encoding/json"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	// "gorm.io/driver/postgres"
	// "gorm.io/gorm"
)

type Product struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"size:255;not null"`
	Price    int
	Quantity int
}
type User struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// func getUsersTest(w http.ResponseWriter, r *http.Request) {
// 	// 🔹 DSN (bazaga ulanish)
// 	dsn := "host=127.0.0.1 user=postgres password=123456 dbname=go_vs_php port=5432 sslmode=disable TimeZone=Asia/Tashkent"

// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		http.Error(w, "❌ Bazaga ulanishda xatolik: "+err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	// 🔹 Query paramlarni olish
// 	pageStr := r.URL.Query().Get("page")
// 	limitStr := r.URL.Query().Get("limit")
// 	search := r.URL.Query().Get("search")

// 	if pageStr == "" {
// 		pageStr = "1"
// 	}
// 	if limitStr == "" {
// 		limitStr = "10"
// 	}

// 	page, err1 := strconv.Atoi(pageStr)
// 	limit, err2 := strconv.Atoi(limitStr)

// 	if err1 != nil || err2 != nil || page <= 0 || limit <= 0 {
// 		http.Error(w, "Invalid page or limit", http.StatusBadRequest)
// 		return
// 	}

// 	offset := (page - 1) * limit

// 	// 🔹 Query yaratish
// 	query := db.Model(&User{})

// 	// 🔹 Agar search bo‘lsa, name yoki email bo‘yicha qidirish
// 	if search != "" {
// 		query = query.Where("name ILIKE ? OR email ILIKE ?", "%"+search+"%", "%"+search+"%")
// 	}

// 	// 🔹 Umumiy sonni sanash
// 	var total int64
// 	query.Count(&total)

// 	// 🔹 Ma’lumotlarni olish — ORDER BY id DESC + LIMIT/OFFSET
// 	var users []User
// 	query.Order("id DESC").Offset(offset).Limit(limit).Find(&users)

// 	// 🔹 JSON javob tayyorlash
// 	response := map[string]interface{}{
// 		"data":  users,
// 		"page":  page,
// 		"limit": limit,
// 		"total": total,
// 		"pages": int(math.Ceil(float64(total) / float64(limit))),
// 	}

//		w.Header().Set("Content-Type", "application/json")
//		json.NewEncoder(w).Encode(response)
//	}
func getUsersTest(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// 🔹 DSN
	dsn := "postgres://postgres:123456@127.0.0.1:5432/go_vs_php?sslmode=disable"

	// 🔹 Pool sozlash
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		http.Error(w, "Config xatosi: "+err.Error(), http.StatusInternalServerError)
		return
	}
	config.MaxConns = 50
	config.MinConns = 5
	config.MaxConnLifetime = time.Hour
	config.HealthCheckPeriod = time.Minute

	// 🔹 Connection pool yaratish
	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		http.Error(w, "Bazaga ulanishda xatolik: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer pool.Close()

	// 🔹 Query paramlar
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")
	search := r.URL.Query().Get("search")

	if pageStr == "" {
		pageStr = "1"
	}
	if limitStr == "" {
		limitStr = "10"
	}

	page, err1 := strconv.Atoi(pageStr)
	limit, err2 := strconv.Atoi(limitStr)
	if err1 != nil || err2 != nil || page <= 0 || limit <= 0 {
		http.Error(w, "Invalid page or limit", http.StatusBadRequest)
		return
	}

	offset := (page - 1) * limit

	// 🔹 SQL query’lar
	baseQuery := `SELECT id, name, email FROM users`
	countQuery := `SELECT COUNT(*) FROM users`

	args := []interface{}{}
	searchCondition := ""

	if search != "" {
		searchCondition = " WHERE name ILIKE $1 OR email ILIKE $2"
		args = append(args, "%"+search+"%", "%"+search+"%")
	}

	// 🔹 Sanash (total)
	var total int64
	err = pool.QueryRow(ctx, countQuery+searchCondition, args...).Scan(&total)
	if err != nil {
		http.Error(w, "Count error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// 🔹 Ma’lumotlarni olish
	query := baseQuery + searchCondition + " ORDER BY id DESC LIMIT $3 OFFSET $4"
	args = append(args, limit, offset)

	rows, err := pool.Query(ctx, query, args...)
	if err != nil {
		http.Error(w, "Query error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	users := make([]User, 0, limit)
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			http.Error(w, "Row scan error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, u)
	}

	if rows.Err() != nil {
		http.Error(w, "Row iteration error: "+rows.Err().Error(), http.StatusInternalServerError)
		return
	}

	// 🔹 JSON response
	response := map[string]interface{}{
		"data":  users,
		"page":  page,
		"limit": limit,
		"total": total,
		"pages": int(math.Ceil(float64(total) / float64(limit))),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
func main() {
	http.HandleFunc("/", getUsersTest)
	http.ListenAndServe(":8080", nil)

	// dsn := "host=127.0.0.1 user=postgres password=123456 dbname=go_data port=5432 sslmode=disable TimeZone=Asia/Tashkent"

	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	panic("❌ Bazaga ulanishda xatolik: " + err.Error())
	// }
	// fmt.Println("✅ PostgreSQL bilan bog‘lanish muvaffaqiyatli!")
	// fmt.Println("✅ Ulanish OK")
	// // db.AutoMigrate(&Product{})
	// fmt.Println("📦 Jadval avtomatik migratsiya qilindi!")

	// db.Create(&Product{Name: "Kofta", Price: 120000, Quantity: 5})
	// var products []Product
	// db.Find(&products)
	// fmt.Println(products)
	// var p Product
	// db.First(&p, 1) // ID=1
	// fmt.Println(p.Name)

	// var products []Product
	// db.Where("price > ?", 100000).Find(&products)

	// db.Model(&p).Update("Price", 150000)

	// db.Delete(&p)

}

// | Funksiya        | Tavsif                           | Misol                                    |
// | --------------- | -------------------------------- | ---------------------------------------- |
// | `AutoMigrate()` | Jadvalni avtomatik yaratadi      | `db.AutoMigrate(&User{})`                |
// | `Model()`       | Jadval bilan ishlash             | `db.Model(&User{})`                      |
// | `Select()`      | Ustunlarni tanlash               | `db.Select("name", "price")`             |
// | `Order()`       | Saralash                         | `db.Order("price desc").Find(&products)` |
// | `Limit()`       | Limit                            | `db.Limit(10).Find(&products)`           |
// | `Joins()`       | SQL JOIN                         | `db.Joins("Category").Find(&products)`   |
// | `Preload()`     | Eager load (Laravel’da `with()`) | `db.Preload("Category").Find(&products)` |

// | Amal                   | Funksiya                | SQL natija       |
// | ---------------------- | ----------------------- | ---------------- |
// | `Find(&x)`             | Barcha yozuvlarni olish | SELECT *         |
// | `First(&x)`            | 1-ta yozuv              | SELECT * LIMIT 1 |
// | `Where("id = ?", 1)`   | Shartli filter          | WHERE id=1       |
// | `Or()`                 | Yoki                    | OR               |
// | `Not()`                | Teskari shart           | WHERE NOT ...    |
// | `Order("price desc")`  | Saralash                | ORDER BY         |
// | `Limit(5)`             | Limit                   | LIMIT 5          |
// | `Offset(10)`           | Offset                  | OFFSET 10        |
// | `Select("id,name")`    | Ustunlarni tanlash      | SELECT id,name   |
// | `Count(&c)`            | Hisoblash               | COUNT(*)         |
// | `Group("category_id")` | Guruhlash               | GROUP BY         |
// | `Preload("Category")`  | Eager loading           | JOIN category    |

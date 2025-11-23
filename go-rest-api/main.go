package main

import (
	"log"
	"math"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	_ "modernc.org/sqlite" // ✅ CGO talab qilmaydigan drayver
)

// MODELLAR

type User struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	// Posts []Post `json:"posts"` // hasMany
}

type Post struct {
	ID      uint   `gorm:"primaryKey" json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  uint   `json:"user_id"`
	User    User   `json:"user"` // belongsTo
}

var db *gorm.DB

// DATABASE ULANISHI
func initDatabase() {
	// Fayl nomi
	dbFile := "app.db"

	// Agar database fayli mavjud bo‘lmasa — avtomatik yaratadi
	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		log.Println("📂 Database fayli topilmadi, yangi fayl yaratilmoqda...")
		file, err := os.Create(dbFile)
		if err != nil {
			log.Fatal("❌ Fayl yaratishda xato:", err)
		}
		file.Close()
	}

	// Ulanish
	var err error
	db, err = gorm.Open(sqlite.Open(dbFile), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Database ulanmadi:", err)
	}
	log.Println("✅ Database ulandi")

	// Migratsiya — jadval bo‘lmasa avtomatik yaratadi
	// err = db.AutoMigrate(&User{}, &Post{})
	if err != nil {
		log.Fatal("❌ Jadval yaratishda xato:", err)
	}
	log.Println("✅ Migratsiyalar bajarildi")
}

// MAIN
func main() {
	initDatabase()
	r := gin.Default()

	// ROUTES
	r.GET("/users-test", getUsersTest)
	r.GET("/users", getUsers)
	r.POST("/users", createUser)
	r.GET("/users/:id", getUserByID)

	r.GET("/posts", getPosts)
	r.POST("/posts", createPost)
	r.GET("/posts/:id", getPostByID)

	log.Println("🚀 Server http://localhost:8080 da ishga tushdi")
	r.Run(":8080")
}

// HANDLERLAR

// -------- USERS ----------
func getUsers(c *gin.Context) {
	var users []User
	db.Preload("Posts").Find(&users)
	c.JSON(http.StatusOK, users)
}
func getUsersTest(c *gin.Context) {
	var users []User

	// Querydan page va limit qiymatlarini olish
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, err1 := strconv.Atoi(pageStr)
	limit, err2 := strconv.Atoi(limitStr)

	if err1 != nil || err2 != nil || page <= 0 || limit <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page or limit"})
		return
	}

	// Qancha o‘tkazib yuborish kerakligini hisoblash
	offset := (page - 1) * limit

	// Umumiy user sonini olish (frontendda sahifalar sonini ko‘rsatish uchun foydali)
	var total int64
	db.Model(&User{}).Count(&total)

	// Ma’lumotlarni limit bilan olish
	db.Offset(offset).Limit(limit).Find(&users)

	// Javob qaytarish
	c.JSON(http.StatusOK, gin.H{
		"data":  users,
		"page":  page,
		"limit": limit,
		"total": total,
		"pages": int(math.Ceil(float64(total) / float64(limit))),
	})
}

func createUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&user)
	c.JSON(http.StatusCreated, user)
}

func getUserByID(c *gin.Context) {
	id := c.Param("id")
	var user User
	if err := db.Preload("Posts").First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User topilmadi"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// -------- POSTS ----------
func getPosts(c *gin.Context) {
	var posts []Post
	db.Preload("User").Find(&posts)
	c.JSON(http.StatusOK, posts)
}

func createPost(c *gin.Context) {
	var post Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&post)
	c.JSON(http.StatusCreated, post)
}

func getPostByID(c *gin.Context) {
	id := c.Param("id")
	var post Post
	if err := db.Preload("User").First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post topilmadi"})
		return
	}
	c.JSON(http.StatusOK, post)
}

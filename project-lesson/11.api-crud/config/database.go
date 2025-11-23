package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "host=localhost user=postgres password=123456 dbname=go-crud port=5432 sslmode=disable TimeZone=Asia/Tashkent"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("❌ Databasega ulanishda xato: " + err.Error())
	}
	DB = db
	fmt.Println("✅ PostgreSQL bilan ulanish muvaffaqiyatli!")
}

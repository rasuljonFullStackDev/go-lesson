package commands

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func MakeModel(name string, withMigration bool) {
	modelDir := "models"
	migDir := "migrations"

	os.MkdirAll(modelDir, 0755)
	os.MkdirAll(migDir, 0755)

	// 🧱 Model faylini yaratish
	modelName := strings.Title(name)
	modelFile := fmt.Sprintf("%s/%s.go", modelDir, strings.ToLower(name))

modelTemplate := fmt.Sprintf(`package models

type %s struct {
	ID        uint   "gorm:\"primaryKey\""
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
`, modelName)

	os.WriteFile(modelFile, []byte(modelTemplate), 0644)
	fmt.Println("✅ Model yaratildi:", modelFile)

	// 🧾 Agar -m flag berilgan bo‘lsa — migratsiya yaratish
	if withMigration {
		timestamp := time.Now().Format("20060102150405")
		migFile := fmt.Sprintf("%s/%s_create_%ss_table.sql", migDir, timestamp, strings.ToLower(name))

		migTemplate := fmt.Sprintf(`-- +goose Up
CREATE TABLE %ss (
	id SERIAL PRIMARY KEY,
	name VARCHAR(255),
	created_at TIMESTAMP DEFAULT NOW(),
	updated_at TIMESTAMP DEFAULT NOW()
);

-- +goose Down
DROP TABLE %ss;
`, strings.ToLower(name), strings.ToLower(name))

		os.WriteFile(migFile, []byte(migTemplate), 0644)
		fmt.Println("✅ Migration yaratildi:", migFile)
	}
}

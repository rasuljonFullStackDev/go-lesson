package commands

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func MakeMigration(name, table string) {
	migDir := "migrations"
	os.MkdirAll(migDir, 0755)

	timestamp := time.Now().Format("20060102150405")
	migFile := fmt.Sprintf("%s/%s_%s.sql", migDir, timestamp, strings.ToLower(name))

	var migTemplate string

	// Agar table nomi berilgan bo‘lsa — ALTER TABLE
	if table != "" {
		migTemplate = fmt.Sprintf(`-- +goose Up
ALTER TABLE %s
ADD COLUMN new_column_name VARCHAR(255);

-- +goose Down
ALTER TABLE %s
DROP COLUMN new_column_name;
`, table, table)
	} else {
		// Yangi jadval yaratish
		migTemplate = fmt.Sprintf(`-- +goose Up
CREATE TABLE %s (
	id SERIAL PRIMARY KEY,
	name VARCHAR(255),
	created_at TIMESTAMP DEFAULT NOW(),
	updated_at TIMESTAMP DEFAULT NOW()
);

-- +goose Down
DROP TABLE %s;
`, strings.ToLower(strings.TrimPrefix(name, "create_")), strings.ToLower(strings.TrimPrefix(name, "create_")))
	}

	os.WriteFile(migFile, []byte(migTemplate), 0644)
	fmt.Println("✅ Migration yaratildi:", migFile)
}

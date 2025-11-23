package commands

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	_ "github.com/lib/pq"
)

func RunMigrate() {
	db, err := sql.Open("postgres", "user=postgres password=123456 dbname=go-crud sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	files, err := ioutil.ReadDir("migrations")
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".sql" {
			content, _ := ioutil.ReadFile("migrations/" + file.Name())
			_, err := db.Exec(string(content))
			if err != nil {
				fmt.Println("❌ Xato:", file.Name(), "-", err)
			} else {
				fmt.Println("✅ Migratsiya bajarildi:", file.Name())
			}
		}
	}
}

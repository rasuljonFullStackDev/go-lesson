package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"11.api-crud/commands"
	"11.api-crud/config"
	"11.api-crud/models"
	"11.api-crud/router"
)

func main() {
	config.ConnectDB()

	config.DB.AutoMigrate(&models.User{}, &models.Category{}, &models.Product{})

	r := router.RegisterRoutes()

	fmt.Println("🚀 Server: http://localhost:8080")
	http.ListenAndServe(":8080", r)

	if len(os.Args) < 2 {
		fmt.Println("⚠️  Buyruq kiriting: make:model, make:migration, migrate, rollback")
		return
	}

	cmd := os.Args[1]

	switch cmd {
	case "make:model":
		name := os.Args[2]
		withMigration := len(os.Args) > 3 && (os.Args[3] == "-m" || os.Args[3] == "--migration")
		commands.MakeModel(name, withMigration)

	case "make:migration":
		name := os.Args[2]
		table := ""
		for _, arg := range os.Args {
			if strings.HasPrefix(arg, "--table=") {
				table = strings.Split(arg, "=")[1]
			}
		}
		commands.MakeMigration(name, table)

	case "migrate":
		commands.RunMigrate()

	case "rollback":
		commands.RunRollback()

	default:
		fmt.Println("❌ Noma’lum buyruq:", cmd)
	}
}

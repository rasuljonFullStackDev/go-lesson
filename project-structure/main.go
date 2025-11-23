package main

import (
	// "./"
	"project-structure/config/database"
	"project-structure/project-structure/routes"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load(".env")

	db := database.Connect()
	routes.RunServer(db)
}

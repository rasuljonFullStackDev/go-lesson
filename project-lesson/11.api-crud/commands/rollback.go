package commands

import (
	"fmt"
	"os/exec"
)

func RunRollback() {
	fmt.Println("⏪ Oxirgi migratsiyani orqaga qaytaryapman...")
	cmd := exec.Command("goose", "postgres", "user=postgres password=123456 dbname=go-crud sslmode=disable", "down")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("❌ Xato:", err)
	}
	fmt.Println(string(out))
}

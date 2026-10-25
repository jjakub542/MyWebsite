package main

import (
	"fmt"
	"os"

	"my-website-backend/internal/config"
	"my-website-backend/internal/database"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run cmd/database/main.go [command]")
		fmt.Println("Available commands:")
		fmt.Println("  clear          - Deletes all data from the database (keeps tables)")
		fmt.Println("  make-superuser - Creates the default admin account")
		os.Exit(1)
	}

	command := os.Args[1]

	cfg := config.LoadConfig()
	db := database.Connect(cfg)

	switch command {
	case "clear":
		fmt.Println("Clearing database data...")
		if err := database.ClearData(db); err != nil {
			fmt.Printf("❌ Failed to clear database data: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("✅ Database cleared successfully. Tables are empty.")

	case "make-superuser":
		fmt.Println("Creating superuser...")
		database.MakeSuperuser(db, "jjakub542@gmail.com", "123")
		fmt.Println("✅ Superuser command executed.")

	default:
		fmt.Printf("Unknown command: %s\n", command)
		os.Exit(1)
	}
}

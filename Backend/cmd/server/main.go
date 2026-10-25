package main

import (
	"fmt"
	"my-website-backend/internal/app"
)

func main() {
	server := app.NewServer()
	fmt.Println("✅ Server is listening on port :8080")
	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}

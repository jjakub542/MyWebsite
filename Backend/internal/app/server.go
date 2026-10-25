package app

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"gorm.io/gorm"

	"my-website-backend/internal/config"
	"my-website-backend/internal/database"
	"my-website-backend/internal/session"
)

type Server struct {
	port  int
	store *session.Store
	db    *gorm.DB
}

func NewServer() *http.Server {

	cfg := config.LoadConfig()
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	db := database.Connect(cfg)
	store := session.NewStore()
	NewServer := &Server{
		port:  port,
		store: store,
		db:    db,
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.Router(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}

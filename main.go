package main

import (
	"auth/database"
	"auth/handler"
	"auth/server"
	"log"
	"net/http"
	"os"
)

var (
	port         = "8081"
	databaseFile = "./sqliteDB/authenticator.db"
)

func init() {
	if env := os.Getenv("PORT"); env != "" {
		port = env
	}
	if env := os.Getenv("DATABASE_FILE_NAME"); env != "" {
		databaseFile = env
	}
}

func main() {
	db, err := database.InitDB(databaseFile)
	if err != nil {
		panic(err)
	}
	server := server.Server{
		Router: http.NewServeMux(),
	}
	h := handler.Handler{
		DB: db,
	}
	server.InitRoute(&h)
	log.Println("started authentication microservice...")
	log.Fatal(http.ListenAndServe(`:`+port, server.Router))

}

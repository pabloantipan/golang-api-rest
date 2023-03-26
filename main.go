package main

import (
	"context"
	"go/golang-api-rest/handlers"
	"go/golang-api-rest/server"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")
	JWT_SECRET := os.Getenv("JWT_SECRET")
	DATABASE_URL := os.Getenv("DATABASE_URL")

	server, err := server.NewServer(context.Background(), &server.Config{
		Port:        PORT,
		JWTSecret:   JWT_SECRET,
		DatabaseUrl: DATABASE_URL,
	})

	if err != nil {
		log.Fatal(err)
	}

	server.Start(BindRoutes)

}

func BindRoutes(server server.Server, router *mux.Router) {
	router.HandleFunc("/", handlers.HomeHandler(server)).Methods((http.MethodGet))
}

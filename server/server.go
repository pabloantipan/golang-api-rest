package server

import (
	"context"
	"errors"
	"go/golang-api-rest/database"
	"go/golang-api-rest/repository"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Config struct {
	Port        string
	JWTSecret   string
	DatabaseUrl string
}

type Server interface {
	Config() *Config
}

type Broker struct {
	config *Config
	router *mux.Router
}

func (broker *Broker) Config() *Config {
	return broker.config
}

func NewServer(context context.Context, config *Config) (*Broker, error) {
	if config.Port == "" {
		return nil, errors.New("port is required")
	}
	if config.JWTSecret == "" {
		return nil, errors.New("jwt secret is required")
	}
	if config.DatabaseUrl == "" {
		return nil, errors.New("database url is required")
	}

	return &Broker{config: config, router: mux.NewRouter()}, nil
}

func (broker *Broker) Start(binder func(server Server, router *mux.Router)) {
	broker.router = mux.NewRouter()
	binder(broker, broker.router)
	repo, err := database.NewPostgresRepository(broker.config.DatabaseUrl)
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}
	repository.SetRepository(repo)

	log.Println("Starting server on port", broker.Config().Port)
	if err := http.ListenAndServe(broker.config.Port, broker.router); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

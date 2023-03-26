package handlers

import (
	"encoding/json"
	"go/golang-api-rest/server"
	"net/http"
)

type HomeResponse struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

func HomeHandler(server server.Server) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/jason")
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(HomeResponse{
			Message: "Welcome to Golang Server",
			Status:  true,
		})
	}
}

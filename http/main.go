package main

import (
	"log"
	"net/http"

	"github.com/Le0nar/golang_practice/http/internal/handler"
	"github.com/Le0nar/golang_practice/http/internal/repository"
	"github.com/Le0nar/golang_practice/http/internal/service"
)

func main() {
	repo := repository.NewRepository()
	service := service.NewService(repo)
	hndlr := handler.NewHandler(service)

	router := hndlr.InitRouter()

	err := http.ListenAndServe(":3000", router)
	if err != nil {
		log.Fatal(err)
	}
}

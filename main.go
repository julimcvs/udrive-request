package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"udrive-request/configs"
	"udrive-request/handler"
)

func main() {
	err := configs.Load()
	api := configs.GetAPI()
	fmt.Printf("API Value: %s", api)
	if err != nil {
		fmt.Printf("Erro ao inicializar API")
		panic(err)
	}

	router := chi.NewRouter()
	router.Post("/request", handler.Create)

	err = http.ListenAndServe(fmt.Sprintf(":%s", configs.GetAPI()), router)
	if err != nil {
		fmt.Printf("Erro ao inicializar API")
		panic(err)
	}
	fmt.Printf("Listening on %s", configs.GetAPI())
}

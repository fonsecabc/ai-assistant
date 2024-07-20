package main

import (
	"ai-assistant/main/configs"
	rest "ai-assistant/main/configs/rest"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()

	rest.PrepareMiddlewares(r)
	rest.PrepareRoutes(r)

	println("Server is running on port " + configs.Env.ServerPort)
	http.ListenAndServe(configs.Env.ServerPort, r)
}

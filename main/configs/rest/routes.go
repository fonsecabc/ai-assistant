package configs

import (
	factories "ai-assistant/main/factories/handlers"
	"net/http"

	"github.com/go-chi/chi"
)

type Handler func(w http.ResponseWriter, r *http.Request)

type Route struct {
	Path    string
	Method  string
	Handler Handler
}

var routes []Route = []Route{
	{
		Path:    "/messenger/webhook",
		Method:  http.MethodPost,
		Handler: factories.NewMessengerWebhookHandlerFactory().Handle,
	},
}

func PrepareRoutes(r *chi.Mux) {
	for _, route := range routes {
		r.Method(route.Method, route.Path, http.HandlerFunc(route.Handler))
	}
}

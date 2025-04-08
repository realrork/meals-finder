package routes

import (
	"net/http"

	"github.com/miloszbo/meals-finder/internal/handlers"
)

func SetupRoutes() http.Handler {
	mux := http.NewServeMux()

	helloWorldHandler := &handlers.HelloWorldHandler{}

	mux.HandleFunc("GET /", helloWorldHandler.Greetings)

	return mux
}


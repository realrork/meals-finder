package server

import (
	"net/http"

	"github.com/miloszbo/meals-finder/internal/handlers"
	"github.com/miloszbo/meals-finder/internal/middlewares"
)

func SetupRoutes() http.Handler {
	mux := http.NewServeMux()

	conn := NewConnection()

	helloWorldHandler := handlers.NewHelloWorldHandler(conn)

	mux.HandleFunc("GET /", helloWorldHandler.Greetings)

	return middlewares.Logging(mux)
}

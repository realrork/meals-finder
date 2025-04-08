package handlers

import (
	"net/http"

	"github.com/miloszbo/meals-finder/internal/services"
)

type HelloWorldHandler struct {
	HelloWorldService services.HelloWorldService
}

func (hwh *HelloWorldHandler) Greetings(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(hwh.HelloWorldService.Queries.SayHelloWorld()))
}


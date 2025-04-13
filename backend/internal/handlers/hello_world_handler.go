package handlers

import (
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/miloszbo/meals-finder/internal/services"
)

type HelloWorldHandler struct {
	HelloWorldService services.HelloWorldService
}

func NewHelloWorldHandler(conn *pgx.Conn) HelloWorldHandler {
	return HelloWorldHandler{
		HelloWorldService: services.NewHelloWorldService(conn),
	}
}

func (hwh *HelloWorldHandler) Greetings(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(hwh.HelloWorldService.SayHelloWorld()))
}

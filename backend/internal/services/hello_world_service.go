package services

import "github.com/jackc/pgx/v5"

type HelloWorldService struct {
	DbConn *pgx.Conn
}

func NewHelloWorldService(conn *pgx.Conn) HelloWorldService {
	return HelloWorldService{
		DbConn: conn,
	}
}

func (hws *HelloWorldService) SayHelloWorld() string {
	return "Hello World"
}

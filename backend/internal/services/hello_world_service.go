package services

type HelloWorldService struct {
	Queries  Queries
	Commands Commands
}

type Queries struct{}

type Commands struct{}

func (q *Queries) SayHelloWorld() string {
	return "Hello World"
}

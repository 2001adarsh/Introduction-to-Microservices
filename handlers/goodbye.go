package handlers

import (
	"log"
	"net/http"
)

type GoodBye struct {
	l *log.Logger
}

func NewGoodBye(l *log.Logger) *GoodBye{
	return &GoodBye{l}
}

func (goodbye *GoodBye) ServeHTTP(writer http.ResponseWriter, request *http.Request)  {
	writer.Write([]byte("Good Bye"))
}
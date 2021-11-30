package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello  {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	h.l.Println("HELLO WORLD!")
	data, err := ioutil.ReadAll(request.Body)
	if err != nil {
		http.Error(writer, "Oops", http.StatusBadRequest)
		return
	}
	defer request.Body.Close()
	h.l.Printf("logged: Hello %s\n", data)
	fmt.Fprintf(writer, "returned to user: Hello %s", data)
}

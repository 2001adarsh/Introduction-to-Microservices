package main

import (
	"context"
	"github.com/Intro_to_Microservices/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	logOp := log.New(os.Stdout, "product-api", log.LstdFlags)

	//creating the handlers
	helloHandler := handlers.NewHello(logOp)
	goodbyeHandler := handlers.NewGoodBye(logOp)

	//creating a new serveMux and register the handlers
	serveMux := http.NewServeMux()
	serveMux.Handle("/", helloHandler)
	serveMux.Handle("/goodbye", goodbyeHandler)

	customServer := &http.Server{
		Addr:         ":9090",           // configure bind address
		Handler:      serveMux,          // set the default handler
		ErrorLog:     logOp,             //set the logger for server
		ReadTimeout:  1 * time.Second,   // max time to read request to the client
		WriteTimeout: 1 * time.Second,   //max time to write a response to the client
		IdleTimeout:  120 * time.Second, //max time for connections using TCP Keep-Alive
	}

	//start the server
	go func() {
		err := customServer.ListenAndServe()
		if err != nil {
			logOp.Fatal(err)
		}
	}()

	// trap sigterm or interrupt and gracefully shutdown the server
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	//Block until a signal is received.
	sig := <-c
	logOp.Println("Got Signal:", sig, ". Hence Gracefully closing down.")

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	cntx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	customServer.Shutdown(cntx)
}

/* Can use local terminal to test these changes using curl commands like:
Remove-item alias:curl -> on Windows powershell.

curl -v http://localhost:9090/
curl -d 'Adarsh Singh' http://localhost:9090/
curl -d '{"name": "Abby Mallard", "original_voice_actor": "Joan Cusack", "animated_debut": "Chicken Little"}' -H "Content-T
ype: application/json" -X POST http://localhost:9090/
*/

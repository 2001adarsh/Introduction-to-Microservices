package main

import (
	"context"
	"github.com/2001adarsh/Intro_to_Microservices/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	logOp := log.New(os.Stdout, "product-api", log.LstdFlags)
	productHandler := handlers.NewProducts(logOp)

	serveMux := mux.NewRouter()
	//serveMux.Handle("/", productHandler)

	getRouter := serveMux.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", productHandler.GetProducts)
	getRouter.Use(productHandler.MiddlewareValidateProduct)

	putRouter := serveMux.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", productHandler.UpdateProduct)
	putRouter.Use(productHandler.MiddlewareValidateProduct)

	postRouter := serveMux.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", productHandler.CreateProduct)
	postRouter.Use(productHandler.MiddlewareValidateProduct)

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

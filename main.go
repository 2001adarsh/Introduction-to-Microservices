package main

import (
	"context"
	"github.com/2001adarsh/Introduction-to-Microservices/data"
	"github.com/2001adarsh/Introduction-to-Microservices/handlers"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	logOp := log.New(os.Stdout, "product-api", log.LstdFlags)
	logOp.Println("[INFO] SERVER STARTUP")

	validator := data.NewValidation()

	//creates handler
	productHandler := handlers.NewProducts(logOp, validator)

	serveMux := mux.NewRouter()

	getRouter := serveMux.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/products", productHandler.ListAll)
	getRouter.HandleFunc("/products/{id:[0-9}+}", productHandler.ListSingle)

	putRouter := serveMux.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/products", productHandler.UpdateProduct)
	putRouter.Use(productHandler.MiddlewareValidateProduct)

	postRouter := serveMux.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/products", productHandler.CreateProduct)
	postRouter.Use(productHandler.MiddlewareValidateProduct)

	deleteRouter := serveMux.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/products/{id:[0-9]+}", productHandler.DeleteProduct)

	// handler for documentation
	opt := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opt, nil)

	getRouter.Handle("/docs", sh)
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

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

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	//HandleFunc: provides a way to specify how requests to a specific route should be handled.
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		log.Println("HELLO WORLD!")
		data, err := ioutil.ReadAll(request.Body)
		if err != nil {
			/*writer.WriteHeader(http.StatusBadRequest)
			writer.Write([]byte("Oops"))*/ //this line can be replaced with http.Error below.
			http.Error(writer, "Oops", http.StatusBadRequest)
			return
		}
		defer request.Body.Close()
		log.Printf("Data %s\n", data)

		//Writing a response back to the user. Using ResponseWriter
		fmt.Fprintf(writer, "Hello %s", data);

	})
	http.HandleFunc("/goodbye", func(writer http.ResponseWriter, request *http.Request) {
		log.Println("GOOD BYE WORLD!")
	})

	// listens on the TCP network address addr and then calls Serve with handler to handle requests on incoming connections.
	http.ListenAndServe(":9090", nil) //if nil is provided in handler, then it processes (default serveMux).

}


/* Can use local terminal to test these changes using curl commands like:
curl -v http://localhost:9090/
curl -d 'Adarsh Singh' http://localhost:9090/
curl -d '{"name": "Abby Mallard", "original_voice_actor": "Joan Cusack", "animated_debut": "Chicken Little"}' -H "Content-T
ype: application/json" -X POST http://localhost:9090/

*/
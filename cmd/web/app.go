package main

import (
	"fmt"
	"net/http"

	"github.com/IsroilMukhitdinov/gomod_test/pkg/handlers"
)

const port = ":9090"

func main() {

	fmt.Println("basic server demo")
	fmt.Println("web application started on port ", port)

	http.HandleFunc("/test", func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(response, "This Is Just A Test")
	})

	http.HandleFunc("/home", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	http.ListenAndServe(port, nil)
}

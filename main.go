package main

import (
	"log"
	"net/http"
	"os"

	"github.com/prestonchoate/product-api/handlers"
)

func main() {
	l := log.New(os.Stdout, "products-api", log.LstdFlags)

	// create the handlers
	baseHandler := handlers.NewBase(l)
	productsHandler := handlers.NewProductsHandler(l)

	// create a new serve mux and register the handlers
	serveMux := http.NewServeMux()
	serveMux.Handle("/", baseHandler)
	serveMux.Handle("/products", productsHandler)

	// Listen for connections on all ip addresses (0.0.0.0)
	// port 9090 
	log.Println("Starting Server")
	err := http.ListenAndServe(":9090", serveMux)
	log.Fatal(err)
}
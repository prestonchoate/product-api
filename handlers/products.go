package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/prestonchoate/product-api/data"
)

type ProductsHandler struct {
	l *log.Logger
}

func NewProductsHandler(l *log.Logger) *ProductsHandler {
	return &ProductsHandler{l}
}

func (productHandler *ProductsHandler) post(responseWriter http.ResponseWriter, request *http.Request) {
	product := data.Product{}
	err := data.FromJson(request.Body, &product)
	if (err != nil) {
		productHandler.l.Println("[ERROR] Deserializing Product", err)
		http.Error(responseWriter, "Error reading product", http.StatusBadRequest)
		return
	}
	product = data.AddProduct(product);
	productHandler.l.Printf("[DEBUG] Inserted product: %#v\n", product)
	responseWriter.Header().Add("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusCreated)
	data.ToJson(responseWriter, product)

}

func (productHandler *ProductsHandler) get(responseWriter http.ResponseWriter, request *http.Request) {
	products := data.GetProducts()
	responseWriter.Header().Add("Content-Type", "application/json")
	err := data.ToJson(responseWriter, &products)
	if (err != nil) {
		productHandler.l.Println("[ERROR] Serializing Products", err)
		fmt.Fprintln(responseWriter, "{\"error\": \"Could not retrieve prodcuts\"")
		return
	}
}

func (productsHandler *ProductsHandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	productsHandler.l.Println("Handling products request")

	switch request.Method {
		case http.MethodGet:
			productsHandler.get(responseWriter, request)
			break
		case http.MethodPost:
			productsHandler.post(responseWriter, request)
		default:
			fmt.Fprintln(responseWriter, "Method not implemented")
	}

}
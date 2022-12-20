package data

import (
	"encoding/json"
	"io"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
}

type Products []*Product

func ToJson(w io.Writer, obj interface{}) error {
	e := json.NewEncoder(w);
	return e.Encode(obj);
}

func FromJson(reader io.Reader, obj interface{}) error {
	decoder := json.NewDecoder(reader)
	return decoder.Decode(obj)
}

func AddProduct(p Product) Product{
	// get the next id in sequence
	maxID := productList[len(productList)-1].ID
	p.ID = maxID + 1
	productList = append(productList, &p)

	return p

}

func GetProducts() Products {
	return productList
}

var productList = [] *Product {
	{
		ID:          1,
		Name:        "Widget A",
		Description: "This is test widget A. It is a tester widget.",
		Price:       12.34,
		SKU:         "widgeta",
	},
	{
		ID:          2,
		Name:        "Widget B",
		Description: "This is test widget B. It is a tester widget",
		Price:       2.29,
		SKU:         "widgetb",
	},
}
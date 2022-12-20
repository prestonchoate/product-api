# product-api
Simple Product API to learn Go

## Specifications
This API should be served on port 9090

Products model is defined in data/product.go. The model should look like:
```
{
  ID:  1,
  Name:  "Widget A",
  Description: "This is test widget A. It is a tester widget.",
  Price:  12.34,
  SKU:  "widgeta",
}
```

Currently only the `GET` and `POST` actions are available to the `/products` endpoint. 

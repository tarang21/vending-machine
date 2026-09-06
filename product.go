package main

type Product struct {
	ProductId   int
	ProductName string
	quantity    int
	cost        int
}

func NewProduct(productId int, productName string, quantity int, cost int) *Product {
	return &Product{productId, productName, quantity, cost}
}

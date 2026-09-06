package main

type Inventory struct {
	Products map[int]*Product
}

func NewInventory() *Inventory {
	return &Inventory{make(map[int]*Product)}
}

func (i *Inventory) AddProduct(p *Product) {
	i.Products[p.ProductId] = p
}

func (i *Inventory) reduceProductQuantity(productId int, quantityToDecrease int) {
	i.Products[productId].quantity =
		i.Products[productId].quantity - quantityToDecrease
}

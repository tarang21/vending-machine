package main

func main() {
	lays := NewProduct(1, "lays", 2, 10)
	pepsi := NewProduct(2, "pepsi", 10, 20)

	vendingMachine := NewVendingMachine()

	vendingMachine.Inventory.AddProduct(lays)
	vendingMachine.Inventory.AddProduct(pepsi)

	// insert money greater than price of product
	vendingMachine.insertMoney(11)
	vendingMachine.selectProduct(*lays)
	vendingMachine.dispensedProduct(*lays)

	// insert money lesser than price of product
	vendingMachine.insertMoney(9)
	vendingMachine.selectProduct(*lays)
	vendingMachine.dispensedProduct(*lays)

	// insert money equal to price of product
	vendingMachine.insertMoney(10)
	vendingMachine.selectProduct(*lays)
	vendingMachine.dispensedProduct(*lays)

	// checking is it is working if all the product quantity has exhaust
	vendingMachine.insertMoney(10)
	vendingMachine.selectProduct(*lays)
	vendingMachine.dispensedProduct(*lays)
}

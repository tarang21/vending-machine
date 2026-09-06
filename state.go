package main

import "fmt"

type State interface {
	insertMoney(money int)
	selectProduct(p Product)
	dispensedProduct(p Product)
	returnChange()
}

type MoneyInsertedState struct {
	VendingMachine *VendingMachine
}

func (m *MoneyInsertedState) insertMoney(amount int) {
	m.VendingMachine.InsertedMoney += amount
	fmt.Println("Amount inserted")
	m.VendingMachine.updateState(&m.VendingMachine.ProductSelectionState)
}

func (m *MoneyInsertedState) selectProduct(p Product) {
	fmt.Println("Insert money to select product")
}

func (m *MoneyInsertedState) dispensedProduct(p Product) {
	fmt.Println("Insert money to select product and dispensed it")
}

func (m *MoneyInsertedState) returnChange() {
	fmt.Println("Money not inserted")
}

type ProductSelectionState struct {
	VendingMachine *VendingMachine
}

func (ps *ProductSelectionState) insertMoney(moneyInserted int) {
	fmt.Println("money is already inserted, select product to proceed")
}

func (ps *ProductSelectionState) selectProduct(p Product) {
	if p.cost > ps.VendingMachine.InsertedMoney {
		fmt.Println("The cost of this product is ", p.cost, ". Please enter correct amount")
		ps.returnChange()
		ps.VendingMachine.updateState(&ps.VendingMachine.MoneyInsertedState)
		return
	}
	if ps.VendingMachine.Inventory.Products[p.ProductId].quantity == 0 {
		fmt.Println("The product is not available")
		ps.returnChange()
		ps.VendingMachine.updateState(&ps.VendingMachine.MoneyInsertedState)
		return
	}

	ps.VendingMachine.ProductSelected = p
	fmt.Println("Product selected is ", p.ProductName)
	ps.VendingMachine.updateState(&ps.VendingMachine.ProductDispensedState)

}

func (ps *ProductSelectionState) dispensedProduct(p Product) {
	fmt.Println("Please first select product")
}

func (ps *ProductSelectionState) returnChange() {
	if ps.VendingMachine.InsertedMoney > 0 {
		fmt.Println("returning all the money", ps.VendingMachine.InsertedMoney)
		ps.VendingMachine.InsertedMoney = 0
	}
}

type ProductDispensedState struct {
	VendingMachine *VendingMachine
}

func (ps *ProductDispensedState) insertMoney(moneyInserted int) {
	fmt.Println("Let first product be dispensed")
}

func (pd *ProductDispensedState) dispensedProduct(p Product) {
	pd.VendingMachine.Inventory.Products[p.ProductId].quantity -= 1
	fmt.Println("the product has been dispatched, the quantity of product left is",
		pd.VendingMachine.Inventory.Products[p.ProductId].quantity)
	pd.VendingMachine.returnChange()
}

func (m *ProductDispensedState) selectProduct(p Product) {
	fmt.Println("The product has already been dispensed")
}

func (ps *ProductDispensedState) returnChange() {
	ps.VendingMachine.ReturnChange = ps.VendingMachine.InsertedMoney - ps.VendingMachine.ProductSelected.cost
	ps.VendingMachine.InsertedMoney = 0
	fmt.Println("Money returned is ", ps.VendingMachine.ReturnChange)
	ps.VendingMachine.updateState(&ps.VendingMachine.MoneyInsertedState)
}

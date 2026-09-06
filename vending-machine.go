package main

type VendingMachine struct {
	Inventory             *Inventory
	State                 State
	MoneyInsertedState    MoneyInsertedState
	ProductSelectionState ProductSelectionState
	ProductDispensedState ProductDispensedState
	InsertedMoney         int
	ReturnChange          int
	ProductSelected       Product
}

func NewVendingMachine() *VendingMachine {
	inventory := NewInventory()
	vm := &VendingMachine{Inventory: inventory}
	vm.MoneyInsertedState = MoneyInsertedState{VendingMachine: vm}
	vm.ProductSelectionState = ProductSelectionState{VendingMachine: vm}
	vm.ProductDispensedState = ProductDispensedState{VendingMachine: vm}
	vm.State = &vm.MoneyInsertedState
	return vm
}

func (vm *VendingMachine) insertMoney(amount int) {
	vm.State.insertMoney(amount)
}

func (vm *VendingMachine) selectProduct(p Product) {
	vm.State.selectProduct(p)
}

func (vm *VendingMachine) dispensedProduct(p Product) {
	vm.State.dispensedProduct(p)
}

func (vm *VendingMachine) returnChange() {
	vm.State.returnChange()
}

func (vm *VendingMachine) updateState(state State) {
	vm.State = state
}

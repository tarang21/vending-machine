# Vending Machine

A small vending machine in Go using the **state design pattern**.

## What it does

- Add and remove products in **inventory**
- Insert money into the vending machine
- Select a product
- Dispense the product and return change

## Types

| Type | Role |
|------|------|
| `Product` | Id, name, quantity, cost |
| `Inventory` | Map of products by id |
| `VendingMachine` | Holds inventory, money, selected product, and current state |
| `State` (interface) | Actions the machine can receive |

### States

Each of these structs **implements the `State` interface** (all four methods: `insertMoney`, `selectProduct`, `dispensedProduct`, `returnChange`). If even one method is missing, Go will not let you assign it to `VendingMachine.State`.

- **MoneyInsertedState** — accept cash, then move to selection
- **ProductSelectionState** — pick a product if there is enough money and stock
- **ProductDispensedState** — vend the item and return change

Actions that do not apply in that state still need a method (print a message). That is how every state stays “linked” to the same interface.

`VendingMachine` forwards each call to whatever `State` is current:

```go
func (vm *VendingMachine) insertMoney(amount int) {
    vm.State.insertMoney(amount)
}
```

## Mental model (state pattern)

```
main  →  VendingMachine  →  State (interface)
                                │
                    ┌───────────┼───────────┐
                    ▼           ▼           ▼
            MoneyInserted  Selection  Dispensed
```

- `main` only talks to `VendingMachine`
- The machine holds **one** current `State`
- **Every** concrete state implements the same `State` interface, so any of them can sit in that field
- The current state runs the method, then may call `updateState` to switch

Typical path: insert money → select product → dispense → return change → insert money again.

Wrong action in a state: print a message, stay put (or refund and go back).

## Why the state pattern helps here

A vending machine is a **sequence of modes**. The same button means different things depending on the mode:

| Action | Money inserted | Product selection | Dispensed |
|--------|----------------|-------------------|-----------|
| Insert money | Add cash, go to selection | “Already inserted” | “Wait until dispensed” |
| Select product | “Insert money first” | Check price & stock | “Already selected” |
| Dispense | Not allowed yet | “Select first” | Reduce stock, return change |

Without states, `insertMoney` / `selectProduct` would fill up with `if current == ...` branches. Every new step (out of stock, refund only, maintenance) would touch those functions again.

With states:

- Each mode is its **own type** with its own rules
- The machine stays **open for new states** (new struct + interface methods) without rewriting a giant switch
- Illegal operations stay in the **wrong** state as no-ops or messages, so you cannot dispense before paying

That is the point of the pattern: **behavior changes with internal state**, without the caller knowing which struct is current.

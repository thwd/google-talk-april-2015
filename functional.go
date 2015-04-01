package main

import (
	"fmt"
)

func subtotal(quantity, unit_price float64) float64 {
	return quantity * unit_price
}

func grand_total(subtotal, discount float64) float64 {
	return subtotal - discount
}

var (
	quantity   = 2.0
	unit_price = 3.20
	discount   = 0.20
)

func main() {
	fmt.Printf("%.2f", grand_total(subtotal(quantity, unit_price), discount))
}

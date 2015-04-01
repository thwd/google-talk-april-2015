package main

import (
	"fmt"
)

func subtotal(quantity, unit_price float64, cont func(float64)) {
	cont(quantity * unit_price)
}

func grand_total(subtotal, discount float64, cont func(float64)) {
	cont(subtotal - discount)
}

var (
	quantity   = 2.0
	unit_price = 3.20
	discount   = 0.20
)

func main() {
	subtotal(quantity, unit_price, func(f float64) {
		grand_total(f, discount, func(f float64) {
			fmt.Printf("%.2f", f)
		})
	})
}

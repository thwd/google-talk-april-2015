package main

import (
	"fmt"
)

func subtotal(quantity, unit_price, discount float64) {
	grand_total(quantity*unit_price, discount)
}

func grand_total(subtotal, discount float64) {
	fmt.Printf("%.2f", subtotal-discount)
}

var (
	quantity   = 2.0
	unit_price = 3.20
	discount   = 0.20
)

func main() {
	subtotal(quantity, unit_price, discount)
}

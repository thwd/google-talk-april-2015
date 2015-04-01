package main

import (
	"fmt"
)

func NewSubtotal(cont func(float64)) func(float64, float64) {
	return func(quantity, unit_price float64) {
		cont(quantity * unit_price)
	}
}

func NewGrandTotal(cont func(float64)) func(float64, float64) {
	return func(subtotal, discount float64) {
		cont(subtotal - discount)
	}
}

var (
	quantity   = 2.0
	unit_price = 3.20
	discount   = 0.20
)

func main() {
	p := func(f float64) {
		fmt.Printf("%.2f", f)
	}

	g := NewGrandTotal(p)

	c := func(x float64) { // HL
		g(x, discount) // HL
	} // HL

	s := NewSubtotal(c)

	s(quantity, unit_price)
}

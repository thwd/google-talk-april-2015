package main

import (
	"fmt"
)

func Split(f func(float64, float64)) (func(float64), func(float64)) {

	var x, y *float64 = nil, nil

	a := func(v float64) {
		x = &v
		if y != nil {
			f(*x, *y)
		}
	}

	b := func(v float64) {
		y = &v
		if x != nil {
			f(*x, *y)
		}
	}

	return a, b
}

func NewSubtotal(cont func(float64)) (func(float64), func(float64)) {
	return Split(func(quantity, unit_price float64) {
		cont(quantity * unit_price)
	})
}

func NewGrandTotal(cont func(float64)) (func(float64), func(float64)) {
	return Split(func(subtotal, discount float64) {
		cont(subtotal - discount)
	})
}

func NewFigure(cont func(float64)) func(float64) {
	return func(f float64) {
		cont(f)
	}
}

func main() {

	p := func(f float64) {
		fmt.Printf("%.2f", f)
	}

	gs, gd := NewGrandTotal(p)
	sq, su := NewSubtotal(gs)

	quantity := NewFigure(sq)
	unit_price := NewFigure(su)
	discount := NewFigure(gd)

	quantity(2.0)
	unit_price(3.20)
	discount(0.20)

}

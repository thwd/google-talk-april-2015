package main

import (
	"fmt"
)

func DeSync(f func(float64, float64)) (func(float64), func(float64)) {

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

func NewFigure(cont func(float64)) func(float64) {
	return cont
}

func main() {
	p := func(f float64) {
		fmt.Printf("%.2f", f)
	}

	g := NewGrandTotal(p)
	g0, g1 := DeSync(g)

	s := NewSubtotal(g0)
	s0, s1 := DeSync(s)

	q := NewFigure(s0)
	u := NewFigure(s1)
	d := NewFigure(g1)

	// done setting up call graph

	q(2.0)
	u(3.20)
	d(0.20)
}

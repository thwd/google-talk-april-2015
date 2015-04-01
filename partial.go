package main

func NewFigure(cont func(float64)) func(float64) {
    return cont
}

func main() {

	quantity := NewFigure(...)
	unit_price := NewFigure(...)
	discount := NewFigure(...)

	quantity(2.0)
	unit_price(3.20)
	discount(0.20)

}

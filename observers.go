package main

import (
	"fmt"
)

type Observable interface {
	Observe(Observer)
}

type Observer interface {
	Notify(Observable, interface{})
}

// impl

type Broadcaster struct {
	observers []Observer
}

func NewBroadcaster() *Broadcaster {
	return &Broadcaster{}
}

func (b Broadcaster) Notify(o Observable, v interface{}) {
	for _, s := range b.observers {
		s.Notify(o, v)
	}
}

func (b *Broadcaster) Observe(o Observer) {
	b.observers = append(b.observers, o)
}

// help

type LambdaObserver func(Observable, interface{})

func (l LambdaObserver) Notify(o Observable, v interface{}) {
	l(o, v)
}

type AllObservable struct {
	broadcaster *Broadcaster
	buffer      []interface{}
}

func All(os ...Observable) Observable {
	a := AllObservable{
		NewBroadcaster(),
		make([]interface{}, len(os)),
	}
	for i, o := range os {
		o.Observe(LambdaObserver(func(_ Observable, v interface{}) {
			a.buffer[i] = v
			for _, v := range a.buffer {
				if v == nil {
					return
				}
			}
			a.broadcaster.Notify(a, a.buffer)
		}))
	}
	return a
}

func (a AllObservable) Observe(o Observer) {
	a.broadcaster.Observe(o)
}

// TODO: forget this shit, return to "reactive spec"
//       func(os ...func(ot)) func(it)

func main() {
	quantity := NewBroadcaster()
	unit_price := NewBroadcaster()
	discount := NewBroadcaster()

	subtotal := All(quantity, unit_price).Observe(LambdaObserver(func(_ Observable, v interface{}) {
		q := v.([]interface{})[0].(float64)
		u := v.([]interface{})[1].(float64)
		// return q * u
		fmt.Println(q * u)
	}))

	grand_total := All(subtotal, discount).Observe(LambdaObserver(func(_ Observable, v interface{}) {
		s := v.([]interface{})[0].(float64)
		d := v.([]interface{})[1].(float64)
		// return s - d
	}))

	grand_total.Observe(LambdaObserver(func(_ Observable, v interface{}) {
		fmt.Println(v)
	}))

	quantity.Notify(nil, 2.0)
	unit_price.Notify(nil, 3.20)
	discount.Notify(nil, 0.20)
}

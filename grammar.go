package main

import (
	"fmt"
)

func NewDigit(cont func(int, []byte)) func([]byte) {
	return func(b []byte) {
		if len(b) > 0 && b[0] > 47 && b[0] < 58 {
			cont((int)(b[0]-48), b[1:])
		}
	}
}

func NewPlus(cont func(byte, []byte)) func([]byte) {
	return func(b []byte) {
		if len(b) > 0 && b[0] == '+' {
			cont(b[0], b[1:])
		}
	}
}

func DeSequence(f func(int, byte, int, []byte)) (func(int, []byte), func(byte, []byte), func(int, []byte)) {

	a, o := (*int)(nil), (*byte)(nil)

	f0 := func(x int, c []byte) {
		if a == nil {
			a = &x
		}
	}

	f1 := func(s byte, c []byte) {
		if a != nil && o == nil {
			o = &s
		}
	}

	f2 := func(x int, c []byte) {
		if a != nil && o != nil {
			g0, g1 := *a, *o
			a, o = nil, nil
			f(g0, g1, x, c)
		}
	}

	return f0, f1, f2
}

func Or(fs ...func(int, []byte)) func(x int, b []byte) {
	return func(x int, b []byte) {
		for _, f := range fs {
			f(x, b)
		}
	}
}

/*
   P = A
   A = A '+' N | N
   N = [0-9]
*/

func main() {

	P := func(x int, b []byte) {
		if len(b) == 0 {
			fmt.Println(x)
		}
	}

	var A func(int, []byte)
	var N func(int, []byte)

	a0, a1, a2 := DeSequence(func(x int, s byte, y int, b []byte) {
		A(x+y, b)
	})

	A = Or(P, a0)
	N = Or(a2, A)

	// done building call graph

	t0 := NewPlus(a1)
	t2 := NewDigit(N)

	// done defining terminals

	bs := []byte("2+3+4+5+6") // = 20

	fmt.Println(`input: "2+3+4+5+6"`)

	for i, _ := range bs {
		t0(bs[i:])
		t2(bs[i:])
	}
}

/*
func DeSequence(f func(int, byte, int, []byte)) (func(int, []byte), func(byte, []byte), func(int, []byte)) {

    type frame struct {
        x int
        s byte
    }

    stack := make([]frame, 0)

    f0 := func(x int, b []byte) {
        stack = append(stack, frame{x, 0})
        fmt.Println("st", stack)
    }

    f1 := func(s byte, b []byte) {
        if len(stack) > 0 {
            stack[len(stack)-1].s = s
        }
    }

    f2 := func(x int, b []byte) {
        if len(stack) > 0 {
            frame := stack[len(stack)-1]
            if frame.s != 0 {
                f(frame.x, frame.s, x, b)
                stack = stack[:len(stack)-1]
            }
        }
    }

    return f0, f1, f2
}


func DeSequence(f func(int, byte, int, []byte)) (func(int, []byte), func(byte, []byte), func(int, []byte)) {

    xs := make([]int, 0)
    o := (*byte)(nil)

    f0 := func(x int, c []byte) {
        xs = append(xs, x)
    }

    f1 := func(s byte, c []byte) {
        if o == nil {
            o = &s
        }
    }

    f2 := func(x int, c []byte) {
        if len(xs) > 0 && o != nil {
            ys := make([]int, len(xs), len(xs))
            copy(ys, xs)
            q := *o
            xs = xs[:0]
            o = nil
            for _, y := range ys {
                f(y, q, x, c)
            }
        }
    }

    return f0, f1, f2
}
*/

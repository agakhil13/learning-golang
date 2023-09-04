package main

import (
	"fmt"
)

func main() {
	n := 34
	fmt.Printf("%v, %T\n", n, n)
	var l uint16 = 32
	fmt.Printf("%v, %T\n", l, l)

	a := 10
	b := 5
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	fmt.Println(a & b)
	fmt.Println(a | b)
	fmt.Println(a ^ b)
	fmt.Println(a &^ b)
	fmt.Println(a << 2)
	fmt.Println(a >> 2)
	c := 40
	var d int8 = 100
	fmt.Println(c + int(d))

	var e float64
	e = 44.13
	e = 13.7e12
	e = 2.1e12
	fmt.Printf("%v, %T\n", e, e)
	var f complex64
	var g complex64
	f = 4.43 + 4.4i
	g = 33 + 1.2i
	fmt.Printf("%v, %T\n", f, f)
	fmt.Printf("%v, %T\n", real(f), real(f))
	fmt.Printf("%v, %T\n", imag(f), imag(f))
	fmt.Println(f - g)
	fmt.Println(f + g)
	fmt.Println(f * g)
	fmt.Println(f / g)
	var h complex128 = complex(12, 5)
	fmt.Printf("%v, %T", h, h)

}

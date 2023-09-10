package main
import (
	"fmt"
)
func main()  {
	// a := 23
	// b := &a
	// fmt.Println(a, *b)
	// a = 11
	// fmt.Println(a, *b)
	// *b = 13
	// c := 34
	// b = &c
	// fmt.Println(a, *b, c, b)

	var a int = 32
	var b *int = &a
	fmt.Println(a, *b)
}
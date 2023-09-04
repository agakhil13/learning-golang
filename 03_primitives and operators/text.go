package main

import (
	"fmt"
)
func main() {
	// s := "This is a string."
	// fmt.Printf("%v, %T\n", s,s)
	// fmt.Printf("%v, %T\n", string(s[2]),s[2])
	// b := []byte(s)
	// fmt.Printf("%v, %T\n", b, b)
	var r rune = 'a'
	fmt.Printf("%v, %T\n", string(r),r)
	var i int32 = 99
	fmt.Printf("%v, %T\n", string(i),i)


}

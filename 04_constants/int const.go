package main

import (
	"fmt"
)
const myConst int = 322

func main()  {
	// const i int = 2
	// const myConst int = i
	// myConst = 22 // Error
	const myConst int = 22
	fmt.Printf("%v, %T", myConst, myConst)
}
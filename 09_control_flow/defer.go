package main

import (
	"fmt"
)

func main()  {
	// defer fmt.Println("start")
	// defer fmt.Println("middle")
	// defer fmt.Println("end")

	a := "start"
	defer fmt.Println(a)
	a = "end"
}
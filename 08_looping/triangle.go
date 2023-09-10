package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		for j := 5; j > i; j-- {
			fmt.Printf(" ")
		}
		for k := 1; k <= i*2+1; k++ {
			fmt.Printf("*")
		}

		fmt.Println()
	}
}

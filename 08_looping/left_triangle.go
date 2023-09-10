package main
import (
	"fmt"
)

func main()  {
	for i := 5; i > 0; i-- {
		for j := 1; j < i; j++{
		fmt.Printf("  ")
		}
		for k := i; k <= 5; k++{
		fmt.Printf(" *")
		}

		fmt.Println()
	}
}
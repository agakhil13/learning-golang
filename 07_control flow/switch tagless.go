package main
import (
	"fmt"
)

func main()  {
	i := 99
	switch {
	case i%2==1:
		fmt.Println("Odd")
	case i%2==0:
		fmt.Println("Even")
	default:
		fmt.Println("Not in range")
	}
}
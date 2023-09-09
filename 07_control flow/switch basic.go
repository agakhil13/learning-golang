package main
import (
	"fmt"
)

func main()  {
	switch i := 2+2; i {
	case 1,3,5,7,9:
		fmt.Println("Odd")
	case 0, 2, 4, 6, 8:
		fmt.Println("Even")
	default:
		fmt.Println("Not in range")
	}
}
package main
import (
	"fmt"
)

func main()  {
	i := 8
	switch {
	case i <=10:
		fmt.Println("Number is less than or equal to 10")
		fallthrough
	case i <= 20:
		fmt.Println("Number is less than or equal to 20")
	default:
		fmt.Println("Number is greater than 20")
	}
}
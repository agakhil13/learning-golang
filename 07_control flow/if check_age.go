package main
import (
	"fmt"
)

func main()  {
	age := 11
	if age < 1 || age > 100 {
		fmt.Println("Enter age between 1 - 100")
	} else {
		if age < 18 {
			fmt.Println("You cannot drive!")
		} else {
			fmt.Println("You can Drive!")
		}
	}
}
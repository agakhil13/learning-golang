package main
import (
	"fmt"
)

func main()  {
	var i int = 0
	loop:
		for i < 5 {
			for j := 0; j < i; j++ {
				fmt.Printf("%v * %v = %v\n",i, j,i*j)
				if i*j == 2 {
					break loop
				}
			}
			i++
		}
	fmt.Println(i)
}
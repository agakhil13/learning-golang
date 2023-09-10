package main
import (
	"fmt"
)

func main()  {
	var i int = 0
	for {
		fmt.Println(i)
		i++
		if i == 5 {
			break
		}
	}
	fmt.Println(i)
}
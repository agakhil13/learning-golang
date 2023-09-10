package main
import (
	"fmt"
)

func main()  {
	var i int = 0
	for i < 5 {
	// for ; i < 5; i++{
	// for i := 0; i < 5; i++ {
		fmt.Println(i)
		i++
	}
	fmt.Println(i)
}
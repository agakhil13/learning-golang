package main
import (
	"fmt"
)

func main()  {
	// s := []int{2,4,6,8,4,2,5}
	// for _ ,v := range s {
	// 	fmt.Println(v)
	// }

	s := "This is a test string!"
	for k, v := range s {
		fmt.Println(k, string(v))
	}
}
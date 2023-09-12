package main
import (
	"fmt"
)

func main()  {
	for i := 0; i < 5; i++ {
		func(i int){
			defer fmt.Println(i)
		}(i)
	}
	func() {
		fmt.Println("Hello Go!")
	}()
}
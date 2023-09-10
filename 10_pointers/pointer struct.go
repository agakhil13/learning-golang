package main
import (
	"fmt"
)
func main()  {
	var ms *mystruct
	// ms = &mystruct{foo: 3432}
	ms = new(mystruct)
	(*ms).foo = 242
	fmt.Println((*ms).foo)	
}

type mystruct struct {
	foo int
}
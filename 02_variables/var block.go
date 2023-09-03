package main
import (
	"fmt"
)

var b int = 31
var (
	i int = 34
	str string = "Hello"
	f float32 = 34.56
	g float32 = 34.56
)


func main()  {
	var (
	i int = 35
	str string = "Hello World"
	f float32 = 33
)
	fmt.Println("The value of variable i is =",i)
	fmt.Println("The value of variable str is =",str)
	fmt.Println("The value of variable f is =",f)
}
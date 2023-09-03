package main
import (
	"fmt"
	"strconv"
)

func main()  {
	var i int = 34
	fmt.Println("The value of variable is =",i)
	fmt.Printf("%v, %T\n",i, i)
	// fmt.Println()
	var s string
	s = strconv.Itoa(i)
	fmt.Println("The value of variable is =",s)
	fmt.Printf("%v, %T\n",s, s)

	// var i float32 = 34.34
	// fmt.Println("The value of variable is =",i)
	// fmt.Printf("%v, %T\n",i, i)
	// // fmt.Println()
	// var s int 
	// s = int(i)
	// fmt.Println("The value of variable is =",s)
	// fmt.Printf("%v, %T\n",s, s)
	
	// var i int = 34
	// fmt.Println("The value of variable is =",i)
	// fmt.Printf("%v, %T",i, i)
	// fmt.Println()
	// var s float32 
	// s = float32(i)
	// fmt.Println("The value of variable is =",s)
	// fmt.Printf("%v, %T",s, s)
}
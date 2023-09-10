package main
import (
	"fmt"
)

func main() {
	fmt.Println("Start")
	panic("Something went wrong")
	fmt.Println("End")
}
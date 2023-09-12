package main
import (
	"fmt"
)

func main()  {
	fmt.Println("This is a sample message!")
	sayMessage("This is a function call!")
}

func sayMessage(msg string){
	fmt.Printf("%s\n", msg);
}
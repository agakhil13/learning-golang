package main
import (
	"fmt"
)

func main()  {
	name := "Amit"
	msg := "Hello"
	sayMessage(&msg, &name)
	fmt.Println(name)
}

func sayMessage(msg *string, name *string){
	fmt.Printf("%s %s\n", *msg, *name);
	*name = "Anil"
	fmt.Println(*name)
}
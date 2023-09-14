package main
import (
	"fmt"
	"time"
)

func main()  {
	var msg string= "This is a sample text"
	go func(msg string)  {
		fmt.Printf(msg)
	}(msg)
	msg= "End"
	time.Sleep(100 * time.Millisecond)
}
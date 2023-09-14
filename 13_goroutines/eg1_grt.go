package main
import (
	"fmt"
	"time"
)

func main()  {
	var msg string= "This is a sample text"
	go func()  {
		fmt.Printf(msg)
	}()
	time.Sleep(100 * time.Millisecond)
}
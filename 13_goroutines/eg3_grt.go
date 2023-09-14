package main
import (
	"fmt"
	"sync"
)
var wg = sync.WaitGroup{}
func main()  {
	var msg string= "This is a sample text"
	wg.Add(1)
	go func(msg string)  {
		fmt.Printf(msg)
		wg.Done()
	}(msg)
	msg= "End"
	wg.Wait()
}
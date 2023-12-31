package main
import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}
func main() {
	ch := make(chan int)
	wg.Add(2)
	go func() {
		i := <-ch
		fmt.Println(i)
		ch <- 27
		wg.Done()
	}()
	go func ()  {
		ch <- 42
		j := <- ch
		fmt.Println(j)
		wg.Done()
	}()
	wg.Wait()
}
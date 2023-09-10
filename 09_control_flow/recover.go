package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("start")
	panicker()
	fmt.Println("End")
}
func panicker() {
	fmt.Println("About to Panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error: ", err)
		}
	}()
	panic("Something bad happen")
	fmt.Println("Done Panicking.")
}

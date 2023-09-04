package main

import (
	"fmt"
)

const (
	_ = iota //Ignore first value by assigning blank identifier
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

const (
	x = 1 << iota
	w
	r
)

func main() {
	var filesize float32 = 1024
	fmt.Printf("%.2fKB\n", filesize/KB)
	var permission byte = w | x
	fmt.Printf("%v permission\n", permission)
}
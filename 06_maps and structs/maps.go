package main

import (
	"fmt"
)

func main() {
	color := make(map[string]int)
	color = map[string]int{
		"red":    23,
		"green":  56,
		"orange": 45,
		"blue":   76,
		"brown":  88,
		"black":  35,
		"white":  72,
	}
	// fmt.Println(color)
	// color["yellow"] = 21
	// fmt.Println(color)
	// fmt.Println(color["yellow"])

	// pop, ok := color["yellow"]
	// fmt.Println(pop, ok)
	// pop, ok := color["red"]
	// fmt.Println(pop, ok)
	// fmt.Println(len(color))

	col := color
	delete(col, "orange") // delete key from the map col
	fmt.Println(col)
	fmt.Println(color)
}

package main
import (
	"fmt"
)

func main()  {
	color := map[string]int {
		"red": 1231,
		"yellow": 424,
		"blue": 131,
		"black": 242,
	}
	c := "red"
	if pop, ok := color[c]; ok {
		fmt.Printf("%v color have code: %v", c, pop)
	}
}
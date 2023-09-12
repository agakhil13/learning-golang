package main
import ( 
	"fmt"
)

func main() {
	s := sum(1,2,3,4,5,6)
	fmt.Printf("The sum is: %v\n", *s)
}

func sum(value ...int) *int{
	fmt.Println(value)
	total := 0
	for _, v:= range value{
		total += v
	}
	return &total
}

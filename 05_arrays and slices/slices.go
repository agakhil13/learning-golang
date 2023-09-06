package main
import (
	"fmt"
)

func main() {
	// a := []int{3,4,5,6}
	// b := a
	// b[2] = 7
	// a[0] = 9
	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Printf("%v, %T\n", a,a)
	// fmt.Printf("Length: %v\n", len(a))


	// a := []int{1,2,3,4,5,6,7,8,9,10}
	// b := a[:]
	// c := a[3:]
	// d := a[:6]
	// e := a[3:6]
	// b[1] = 45
	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)
	// fmt.Println(d)
	// fmt.Println(e)


	// a := make([]int, 3 ,35)
	// fmt.Printf("%v\n", a)
	// fmt.Printf("Length: %v\n", len(a))
	// fmt.Printf("Capacity: %v\n", cap(a))


	// a := []int{}
	// fmt.Println(a)
	// fmt.Printf("Length: %v\n", len(a))
	// fmt.Printf("Capacity: %v\n", cap(a))
	// a = append(a, 2)
	// fmt.Println(a)
	// fmt.Printf("Length: %v\n", len(a))
	// fmt.Printf("Capacity: %v\n", cap(a))
	// a = append(a, 1,2,34,5,6,7)
	// fmt.Println(a)
	// fmt.Printf("Length: %v\n", len(a))
	// fmt.Printf("Capacity: %v\n", cap(a))
	// a = append(a, []int{132,3,4,4,5,67,8}...)
	// fmt.Println(a)
	// fmt.Printf("Length: %v\n", len(a))
	// fmt.Printf("Capacity: %v\n", cap(a))


	a := []int{1,2,3,4,5}
	fmt.Println(a)
	b := append(a[:2], a[3:]...)
	a[0] = 5
	fmt.Println(b)
	fmt.Println(a)


}
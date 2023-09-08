package main
import (
	"fmt"
)

type Person struct {
	fname string
	lname string
	age int
	hobby []string
}

func main(){
	p1 := Person{
		fname: "Amit",
		lname: "khanna",
		age: 33,
		hobby: []string{
			"chess",
			"football",
		},
	}
	p2 := Person{
		"Amit",
		"khanna",
		33,
		[]string{
			"chess",
			"football",
		},
	}
	fmt.Println(p1.hobby[1], p2, p2.age)

	aP3 := struct{name string}{name: "Rahul"}
	bP4:= &struct{ name string }{"Rahul"} //pointer to a structure
	bP4.name = "Sonet"
	fmt.Println(aP3, bP4)
}
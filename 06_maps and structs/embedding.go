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

type Student struct {
	Person
	roll_no int
	class string
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

	s1 := Student{}
	s1.fname = "Anuj"
	s1.lname = "Sharma"
	s1.age = 23
	s1.hobby = []string{
		"cricket",
		"singing",
	}
	s1.roll_no = 32
	s1.class = "BCA"


	s2 := Student{
		Person: Person{
			fname: "Asif",
			lname: "Verma",
			age: 21,
			hobby: []string{
				"chess",
				"ludo",
			},
		},
			roll_no: 11,
			class: "MBA",
	}
	fmt.Println(p1.hobby[1])
	fmt.Println(s1, s2)

}
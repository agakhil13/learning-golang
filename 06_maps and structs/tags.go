package main
import (
	"fmt"
	"reflect"
)

type Person struct {
	fname string `required max:"100"`
	lname string
	age int
	hobby []string
}


func main(){
	// p1 := Person{
	// 	fname: "Amit",
	// 	lname: "khanna",
	// 	age: 33,
	// 	hobby: []string{
	// 		"chess",
	// 		"football",
	// 	},
	// }
	// fmt.Println(p1.hobby[1])
	e := reflect.TypeOf(Person{})
	fmt.Println(e)
	for i := 0; i < e.NumField(); i++ {
		varName := e.Field(i).Name
		println(varName)
	}
	field, _ := e.FieldByName("fname")
	fmt.Println(field)

}
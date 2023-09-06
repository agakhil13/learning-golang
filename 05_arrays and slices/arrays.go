package main
import (
	"fmt"
)

func main()  {
	// grade1 := 98
	// grade2 := 44
	// grade3 := 68
	// fmt.Printf("Grades: %v, %v, %v", grade1, grade2, grade3)


	// grades := [3]int{98, 44, 68}
	// grades := [...]int{98, 44, 68}
	// fmt.Printf("Grades: %v",grades)


	// var student [5]string
	// fmt.Printf("Students: %v\n", student)
	// student[0] = "Amit"
	// student[1] = "Anil"
	// student[2] = "Akash"
	// student[3] = "Asif"
	// student[4] = "Abbas"
	// fmt.Printf("Students: %v\n", student)
	// fmt.Printf("Number of students: %v\n", len(student))


	// var identityMatrix [3][3]int
	// // identityMatrix = [3][3]int{{1,0,0}, {0,1,0}, {0,0,1}}
	// identityMatrix[0] = [3]int{1,0,0}
	// identityMatrix[1] = [3]int{0,1,0}
	// identityMatrix[2] = [3]int{0,0,1}
	// fmt.Println(identityMatrix)


	a := [...]int{3,5,6,7}
	// b := a
	b := &a
	b[1] = 4
	fmt.Println(a)
	fmt.Println(b)
}
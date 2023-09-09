package main
import (
	"fmt"
)

func main()  {
	number := 45
	guess := 450

	if guess < 1 || guess > 100 {
		fmt.Printf("You have entered %v, number must be between 1 - 100", guess)
	} /* if guess >= 1 && guess <=100  */ else {
		if guess < number {
			fmt.Println("Too low!")
		}
		if guess > number {
			fmt.Println("Too high!")
		}
		if guess == number{
			fmt.Println("You got it!")
		}
		fmt.Println(guess < number, guess > number, guess == number)
	}
}
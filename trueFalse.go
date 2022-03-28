package main

import "fmt"

func compare(value int) string {
	//do not change this variable resultMessage, secretValue
	resultMessge := ""
	secretValue := 88

	//Insert your code from here
	if value == secretValue {
		resultMessge = "correct"
	}

	if value > secretValue {
		resultMessge = "too high"
		fmt.Println(resultMessge)
	}

	if value < secretValue {
		resultMessge = "too low"
		fmt.Println(resultMessge)
	}

	//do not remove this line
	return resultMessge
}

func main() {
	var guess int
	for {
		fmt.Println("Enter an integer value: ")
		fmt.Scanln(&guess)
		// fmt.Printf("%T, %v\n", guess, guess)
		result := compare(guess)
		if result == "correct" {
			fmt.Println("guess correct")
			break
		}
	}

}

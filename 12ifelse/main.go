package main

import "fmt"

func main() {
	fmt.Println("learna about ifelse")

	loginCount := 20
	var result string

	if loginCount > 10 {
		result = "Regular User"
	} else if loginCount < 10 {
		result = "not regular User"
	} else {
		result = "Exactly 10"
	}
	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	if num := 2; num < 10 {
		fmt.Println("Number is less than 10")
	} else {
		fmt.Println("Number is greater than 10")
	}
}

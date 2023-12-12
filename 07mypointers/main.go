package main

import "fmt"

func main() {

	fmt.Println("welcome to pointers")
	// var ptr *int
	// fmt.Println("The values is ", ptr)
	myNumber := 34
	var ptr = &myNumber
	fmt.Println("The actual address is", ptr)
	fmt.Println("The actual address is", *ptr)
	//fmt.Println("The actual address is", &mynumber)

	*ptr = *ptr + 2
	fmt.Println("The new value is ", *ptr)
}

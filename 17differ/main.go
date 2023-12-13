package main

import "fmt"

func main() {
	defer fmt.Println("hello my name is ")
	defer fmt.Println("Patel")
	defer fmt.Println(" Khushal")

	fmt.Println("Learn differ in go")

	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to User Input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name:")

	//comma ok || err err

	input, _ := reader.ReadString('\n') //read upto \n
	fmt.Println("Thank you", input)
	fmt.Printf("Type of input is %T \n", input)
}

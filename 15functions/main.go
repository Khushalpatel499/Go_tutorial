package main

import "fmt"

func main() {
	fmt.Println("learn function in go")
	greeter()
	result := adder(3, 5)
	fmt.Println("The addition is:", result)

	proRes, myMess := proAdder(3, 4, 5, 1)

	fmt.Println("The new additon is :", proRes)
	fmt.Println("The new additon is :", myMess)
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo

}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "Khushal"
}
func greeter() {
	fmt.Println("hello Khushal")
}

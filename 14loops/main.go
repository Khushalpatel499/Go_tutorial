package main

import "fmt"

func main() {
	fmt.Println("new learn loops in go")
	days := []string{"Monday", "Tuesday", "Wednesday", "Friday", "saturday"}
	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	for input, day := range days {
		fmt.Printf("The value input %v,The value of day:%v\n", input, day)
	}

	roughValue := 5
	for roughValue < 10 {
		if roughValue == 7 {
			goto lco
		}
		if roughValue == 8 {
			break
		}
		fmt.Println(roughValue)
		roughValue++
	}

	//goto
lco:
	fmt.Println("Jump")
}

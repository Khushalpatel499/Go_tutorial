package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("switch cases in golang")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("the dice number is:", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice values is 1 and you can open")
	case 2:
		fmt.Println("You can move 2 step")
	case 3:
		fmt.Println("You can move 3 step")
	case 4:
		fmt.Println("You can move 4 step")
		fallthrough //next case also run after case
	case 5:
		fmt.Println("You can move 5 step")
	case 6:
		fmt.Println("You can move 6 step and roll it again")
	default:
		fmt.Println("Not possible")

	}
}

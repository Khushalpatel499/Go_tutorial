package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our Pizza app")
	fmt.Println("Please rate our Pizza from 1 to 5")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64) //by using input we see that "4\n" come so we cnt convert into number

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("ADD 1 to this rating: ", numRating+1)
	}
}

package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices")
	var fruitList = []string{"cherry", "banana", "orange"}
	fmt.Println("The fruitList is:", fruitList)
	fmt.Printf("The type of fruitList is %T \n", fruitList)

	fruitList = append(fruitList, "apple", "mango")
	fmt.Println("The fruitList is:", fruitList)
	fruitList = append(fruitList[1:])
	fmt.Println("The fruitList is:", fruitList)
	fruitList = append(fruitList[1:3])
	fmt.Println("The fruitList is:", fruitList)
	fruitList = append(fruitList[:3])
	fmt.Println("The fruitList is:", fruitList)

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 946
	highScores[2] = 534
	highScores[3] = 834
	//highScores[4] = 777

	highScores = append(highScores, 556, 444, 898) //   entire memory allocation hapeen agian

	fmt.Println("highScore is:", highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	//how to remove a value based on index in slices
	var courses = []string{"javascript", "ruby", "go", "c++", "java"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}

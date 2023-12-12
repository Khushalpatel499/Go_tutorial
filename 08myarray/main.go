package main

import "fmt"

func main() {
	fmt.Println("Welcome my array in go")

	var fruitList [4]string
	fruitList[0] = "mango"
	fruitList[1] = "cherry"
	fruitList[3] = "banana"
	fmt.Println("The fruit list is", fruitList)
	fmt.Println("The fruit list is", len(fruitList))

	var vegList = [5]string{"beans", "potato", "mashroom"}
	fmt.Println("veg List is :", vegList)
	fmt.Println("veg List is :", len(vegList))
}

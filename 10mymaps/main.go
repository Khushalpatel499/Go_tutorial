package main

import "fmt"

func main() {
	fmt.Println("Maps in golang")
	languages := make(map[string]string)
	languages["Js"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["Py"] = "Python"

	fmt.Println(languages)
	fmt.Println(languages["Js"])
	fmt.Println(languages["Py"])

	delete(languages, "RB")

	fmt.Println(languages)

	//loops are interesting in go

	for key, value := range languages {
		fmt.Printf("For key %v, the values is %v \n", key, value)
	}

	for _, value := range languages {
		fmt.Printf("For key v, the values is %v \n", value)
	}

}

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to the files")
	content := "Hello this is my files for learning"

	file, err := os.Create("./myfile.txt")
	// if err != nil {
	// 	panic(err)
	// }
	checkNilError(err)
	length, err := io.WriteString(file, content)

	checkNilError(err)
	fmt.Println("Length is:", length)
	defer file.Close()
	ReadFile("./myfile.txt")
}

func ReadFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNilError(err)
	fmt.Println("The data inside file is:", databyte)
	fmt.Println("The data inside file is:", string(databyte))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}

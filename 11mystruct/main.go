package main

import "fmt"

func main() {
	fmt.Println("The new chapter struct in go")

	Khushal := User{"khushal", "khushalpatel@gmail.com", true, 18}
	fmt.Println(Khushal)
	fmt.Printf("Detail of khushal is:%+v\n", Khushal)
	fmt.Printf("the name is:%v and email is:%v\n", Khushal.Name, Khushal.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

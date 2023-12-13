package main

import "fmt"

func main() {
	fmt.Println("The new chapter methods in go")

	Khushal := User{"khushal", "khushalpatel@gmail.com", true, 18}
	fmt.Println(Khushal)
	fmt.Printf("Detail of khushal is:%+v\n", Khushal)
	fmt.Printf("the name is:%v and email is:%v\n", Khushal.Name, Khushal.Email)
	Khushal.GetStatus()
	Khushal.NewEmail()
	fmt.Printf("the name is:%v and email is:%v\n", Khushal.Name, Khushal.Email) //here we can see the copy is only passes
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("The status is:", u.Status)
}

func (u User) NewEmail() {
	u.Email = "kh@gamil.com"
	fmt.Println("the new email is:", u.Email)
}

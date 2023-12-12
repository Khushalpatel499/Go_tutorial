package main

import "fmt"

//jwtToken :=10(not allowed)
//var anotherValue = 300(allowed)

const LoginToken string = "xfgddffg" //public

func main() {
	fmt.Println("Variables")
	var username string = "Khushal"
	fmt.Println(username)
	fmt.Printf("Variable is type of:%T and  value is:%v \n", username, username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is type of:%T and  value is:%v \n", isLoggedIn, isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is type of:%T and  value is:%v \n", smallVal, smallVal)

	var smallFloat float32 = 255.4445485940
	fmt.Println(smallVal)
	fmt.Printf("Variable is type of:%T and  value is:%v \n", smallFloat, smallFloat)

	//default values and aliases

	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is type of:%T and  value is:%v \n", anotherVariable, anotherVariable)

	var anotherString string
	fmt.Println(anotherString)
	fmt.Printf("Variable is type of:%T and  value is:%v \n", anotherString, anotherString)

	var isAnotherLogged bool
	fmt.Println(isAnotherLogged)
	fmt.Printf("Variable is type of:%T and  value is:%v \n", isAnotherLogged, isAnotherLogged)

	//implicit type of declaration

	var website = "Khushalpatel.in"
	fmt.Println(website)

	//no var style

	numberOfUser := 5
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)

}

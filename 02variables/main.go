package main

import "fmt"

const LOGINTOKEN string = "ksajdhkahfskashf" // public

func main() {
	var username string = "Abhay"
	fmt.Println(username)
	fmt.Printf("Variable is of type : %T \n", username)

	var isloggedIn bool = true
	fmt.Println(isloggedIn)
	fmt.Printf("Variable is of type : %T \n", isloggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type : %T \n", smallVal)

	var smallFloat float64 = 255.1232363463636
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type : %T \n", smallFloat)

	// default values and some aliases

	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type : %T \n", anotherVariable)

	var anotherstring string
	fmt.Println(anotherstring)
	fmt.Printf("Variable is of type : %T \n", anotherstring)

	//implicit type

	var website = "yahoo.com"
	fmt.Println(website)
	fmt.Printf("Variable is of type : %T \n", website)

	// no var style

	numberOfUser := 20
	fmt.Println(numberOfUser)
	fmt.Printf("Variable is of type : %T \n", numberOfUser)

	fmt.Println(LOGINTOKEN)
	fmt.Printf("Variable is of type : %T \n", LOGINTOKEN)

}

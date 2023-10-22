package main

import "fmt"

func main() {
	fmt.Println("welcome")
	loginCount := 23

	var result string

	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 10 && loginCount < 20 {
		result = "Havy User"
	} else {
		result = "fuk"
	}
	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	if num := 3; num < 5 {
		fmt.Println("num is less than 5")
	} else {
		fmt.Println("Num is greater than 5")
	}

}

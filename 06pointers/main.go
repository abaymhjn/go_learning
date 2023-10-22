package main

import "fmt"

func main() {
	fmt.Println("Wlcome to a class on pointers")

	/* var ptr *int
	fmt.Println("Value of pointer is ", ptr)
	*/

	myNumber := 23

	var ptr = &myNumber

	fmt.Println(ptr)

	myNumber = 25
	// value represent by *pointer
	fmt.Println(*ptr)

	*ptr = *ptr + 2

	fmt.Println(myNumber)
}

package main

import (
	"fmt"
)

/*
	Example explained

In this example there are 6 variables:
student1
student2
x
a
b
c
These variables are declared but they have not been assigned initial values.

By running the code, we can see that they already have the default values of their respective types:

a is ""
b is 0
c is 0
*/
func main() {
	var student1 string = "123"
	var student2 = "abhay"
	x := 4
	var a string
	var b int
	var c float32
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)
	c = 3
	fmt.Println(c)
}

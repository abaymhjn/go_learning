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
var y string

const pi = 3.14

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
	y = "2"
	fmt.Println(y)
	fmt.Println(c)
	//single line multiple valiables define
	var m, n, o = 10, 20, 30
	fmt.Println(m, n, o)

	//define vars as block

	var (
		d   int
		e_d string = "abhay"
	)
	d = 2000
	fmt.Println(d, e_d)
	// constants
	const cont_abhay, cont_1, cont_2 = 12, "abhay", 213
	//cont_abhay = 20
	fmt.Println(cont_abhay, cont_1, cont_2)

	const (
		abc     int    = 234
		sexy    string = "very"
		no_sexy        = 3.5
	)

	fmt.Println(abc, sexy, no_sexy)

	fmt.Println(pi)
	//add space if all variables/constant are not string
	//single '' outout n as variable not new line So use "" for new line
	fmt.Print(abc, no_sexy, "\n")

	fmt.Print(abc, sexy)
}

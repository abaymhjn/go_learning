/*
	The float data types are used to store positive and negative numbers with a decimal point, like 35.3, -2.34, or 3597.34987.

The float data type has two keywords:

Type	Size	Range
float32	32 bits	-3.4e+38 to 3.4e+38.
float64	64 bits	-1.7e+308 to +1.7e+308.
there is no unsigned floating data type
*/
package main

import (
	"fmt"
)

func main() {
	var x float32 = 123.78
	var y float64 = 1.4e+308
	fmt.Printf("x Type: %T, value: %v\n", x, x)
	fmt.Printf("y Type: %T, value: %v\n", y, y)
}

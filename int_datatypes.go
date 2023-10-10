/*
	Signed integers - can store both positive and negative values

Unsigned integers - can only store non-negative values
Type	Size	Range
int	Depends on platform:
32 bits in 32 bit systems and
64 bit in 64 bit systems	-2147483648 to 2147483647 in 32 bit systems and
-9223372036854775808 to 9223372036854775807 in 64 bit systems
int8	8 bits/1 byte	-128 to 127
int16	16 bits/2 byte	-32768 to 32767
int32	32 bits/4 byte	-2147483648 to 2147483647
int64	64 bits/8 byte	-9223372036854775808 to 9223372036854775807
*/
/*
unsigned integers:
Type	Size	Range
uint	Depends on platform:
32 bits in 32 bit systems and
64 bit in 64 bit systems	0 to 4294967295 in 32 bit systems and
0 to 18446744073709551615 in 64 bit systems
uint8	8 bits/1 byte	0 to 255
uint16	16 bits/2 byte	0 to 65535
uint32	32 bits/4 byte	0 to 4294967295
uint64	64 bits/8 byte	0 to 18446744073709551615 */
package main

import (
	"fmt"
)

func main() {
	//Signed integers, declared with one of the int keywords, can store both positive and negative values:
	var x int = 500
	var y int = -4500
	fmt.Printf("x Type: %T, value: %v\n", x, x)
	fmt.Printf("y Type: %T, value: %v\n", y, y)

	//Unsigned integers, declared with one of the uint keywords, can only store non-negative values:
	var a uint = 500
	var b uint = 4500
	fmt.Printf("a Type: %T, value: %v\n", a, a)
	fmt.Printf("b Type: %T, value: %v\n", b, b)
}

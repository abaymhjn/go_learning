package main

import "fmt"

func main() {
	fmt.Println("Welcome")
	greet()
	result := add(3, 5)
	fmt.Println(result)

	proRes := proAdder(2, 3, 4, 5, 6)
	fmt.Println(proRes)

	proRes1, mess := proAdder1(2, 3, 4, 5, 6)
	fmt.Println(proRes1, mess)
}

func greet() {
	fmt.Println("Namaste")
}

func add(x int, y int) int {
	return x + y
}

func proAdder(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}

func proAdder1(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "fuk you"
}

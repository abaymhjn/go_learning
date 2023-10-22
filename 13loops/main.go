package main

import "fmt"

func main() {
	fmt.Println("Welcome")
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	fmt.Println(days)

	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	for i := range days {
		fmt.Println(days[i])
	}

	for index, day := range days {
		fmt.Println("index is ", index, " day is ", day)
	}

	for _, day := range days {
		fmt.Println(" day is ", day)
	}

	rougueValue := 2

	for rougueValue < 10 {
		if rougueValue == 3 {
			goto lco
		}
		if rougueValue == 5 {
			break
		}
		fmt.Println(rougueValue)
		rougueValue++
	}

	for rougueValue < 10 {
		if rougueValue == 5 {
			rougueValue++
			continue
		}
		fmt.Println(rougueValue)
		rougueValue++
	}
lco:
	fmt.Println("Jumped")
}

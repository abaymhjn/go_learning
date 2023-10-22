package main

import "fmt"

func main() {
	fmt.Println("Welcome")

	languages := make(map[string]string)

	languages["JS"] = "javascript"
	languages["PHP"] = "PHP"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println(languages)

	fmt.Println(languages["RB"])

	delete(languages, "RB")
	fmt.Println(languages)

	//loop in map
	for _, value := range languages {
		fmt.Println("Values", value)
	}
}

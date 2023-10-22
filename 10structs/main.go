package main

import "fmt"

func main() {
	fmt.Println("Welcome")
	//no inheritance No super No Parent
	hitesh := User{"Hitesh", "hitesh@go.dev", true, 30}
	fmt.Println(hitesh)

	fmt.Printf("%+v\n", hitesh)

	fmt.Println(hitesh.Name)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

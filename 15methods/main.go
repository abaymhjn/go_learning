package main

import "fmt"

func main() {
	fmt.Println("Welcome")
	hitesh := Users{"histesh", "anbha@adkkkd", 12, true}
	hitesh.GetStatus()
	hitesh.NewMail()
	fmt.Println(hitesh)
}

type Users struct {
	Name   string
	Email  string
	Age    int
	Status bool
}

func (u Users) GetStatus() {
	fmt.Println(u.Status)
}

func (u Users) NewMail() {
	u.Email = "abhay@gmail.com"
	fmt.Println(u.Email)
}

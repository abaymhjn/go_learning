package main

import "fmt"

func main() {
	fmt.Println("welcome to array in golang")

	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Mango"
	fruitList[3] = "Peach"

	fmt.Println(fruitList)
	fmt.Println(len(fruitList))
	fruitList[2] = "banana"
	fmt.Println(fruitList)

	var vegList = [3]string{"potato", "beans", "mus"}
	fmt.Println(vegList)
}

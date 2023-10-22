package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome")
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1

	fmt.Println(diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice 1 and you can open")
	case 2:
		fmt.Println("Dice 2 and you can open")
	case 3:
		fmt.Println("Dice 3 and you can open")
		fallthrough
	case 4:
		fmt.Println("Dice 4 and you can open")
		fallthrough
	case 5:
		fmt.Println("Dice 5 and you can open")
	case 6:
		fmt.Println("Dice 6 and you can open")
		fallthrough
	default:
		fmt.Println("Dice ")
	}
}

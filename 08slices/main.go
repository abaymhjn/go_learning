package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("welcome")
	var fruitList = []string{"banana", "apple"}
	fmt.Printf("Type is %T\n", fruitList)
	fmt.Println(fruitList)
	fruitList = append(fruitList, "mango")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:])
	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867
	//highScores[4] = 777
	highScores = append(highScores, 555, 666, 321)
	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)

	//How to remove value from slices
	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)

	var index int = 2

	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}

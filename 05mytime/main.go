package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome to time study")

	presentTime := time.Now()

	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-05-2006 15:05 Monday"))

	created_date := time.Date(2020, time.April, 10, 23, 23, 0, 0, time.UTC)

	fmt.Println(created_date)
	fmt.Println(created_date.Format("01-02-2006 Monday"))
}

package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome")
	content := "This need to insert in file"

	os.WriteFile("sample.txt", []byte(content), 0666)
	readFile("sample.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNilErr(err)
	fmt.Println(string(databyte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}

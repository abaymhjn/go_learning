package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("Welcome")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	fmt.Printf("Type %T", response)
	//fmt.Println(response)
	defer response.Body.Close()

	res, errs := ioutil.ReadAll(response.Body)
	if errs != nil {
		panic(errs)
	}
	println(string(res))
}

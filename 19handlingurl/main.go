package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://loc.dev:3000/learn?coursename=reactjs&paymentid=12"

func main() {
	fmt.Println("Welcome")
	param, err := url.Parse(myurl)
	if err != nil {
		panic(err)
	}
	qparam := param.Query()

	for index, val := range qparam {
		fmt.Println(index, val)
	}

	parts_of_url := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=abhay",
	}

	another_url := parts_of_url.String()

	fmt.Println(another_url)
}

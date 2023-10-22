package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to webserver")
	perform_get_request()
	perform_post_json_request()
	perform_post_form_data()
}

func perform_get_request() {
	const my_url = "http://localhost:8000/get"
	response, err := http.Get(my_url)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code : ", response.StatusCode)
	fmt.Println("Content Length", response.ContentLength)
	content, errs := io.ReadAll(response.Body)
	if errs != nil {
		panic(errs)
	}
	fmt.Println(string(content))

	var reponse_string strings.Builder

	byteCount, _ := reponse_string.Write(content)
	fmt.Println(byteCount)
	fmt.Println(reponse_string.String())
}

func perform_post_json_request() {
	const my_url = "http://localhost:8000/post"

	// fake json payload

	request_body := strings.NewReader(`
	    {
			"name" : "Abhay",
			"age" : "37"
		}
	`)
	response, err := http.Post(my_url, "application/json", request_body)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, errs := io.ReadAll(response.Body)
	if errs != nil {
		panic(errs)
	}
	fmt.Println(string(content))
}

func perform_post_form_data() {
	const my_url = "http://localhost:8000/postform"

	// formdata
	data := url.Values{}
	data.Add("firstname", "hitesh")
	data.Add("lastname", "mahajan")

	response, err := http.PostForm(my_url, data)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)

	fmt.Println(string(content))
}

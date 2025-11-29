package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("welcome to make web request using go")
	// PerformGetRequest()
	// PerformPostJsonRequest()

	PerformPostFormRequest()
}

func PerformGetRequest() {
	const myUrl = "http://localhost:3000/get"

	res, err := http.Get(myUrl)
	CheckNilErr(err)

	fmt.Println("Status code: ", res.StatusCode)
	fmt.Println("content lenght: ", res.ContentLength)

	content, err := io.ReadAll(res.Body)
	CheckNilErr(err)
	textcontent := string(content)
	fmt.Println(textcontent)

	defer res.Body.Close()

}

func CheckNilErr(err error) {
	if err != nil {
		panic(err)
	}
}

func PerformPostJsonRequest() {
	const myUrl = "http://localhost:3000/post"

	// fake json payload
	requestBody := strings.NewReader(`
		{
			"coursename" : "learn go lang",
			"price" : "1000",
			"platform" : "learncodeonline.in"
		}
	`)

	response, err := http.Post(myUrl, "application/json", requestBody)
	CheckNilErr(err)
	defer response.Body.Close()

	content, err := io.ReadAll(response.Body)
	CheckNilErr(err)
	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const myUrl = "http://localhost:3000/postform"

	// form data
	// url.Values is a map
	data := url.Values{
		"firstname": {"john"},
		"lastname":  {"doe"},
		"email":     {"satvik@gmail.com"},
	}

	data.Add("phoneNo", "1234567890")

	response, err := http.PostForm(myUrl, data)
	CheckNilErr(err)

	content, err := io.ReadAll(response.Body)
	CheckNilErr(err)
	fmt.Println(string(content))	

	defer response.Body.Close()
}

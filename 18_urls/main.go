package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=1234"

func main() {
	fmt.Println("Welcome to url handling")
	fmt.Println("URL is: ", myurl)

	result, _ := url.Parse(myurl)
	fmt.Println("Scheme: ", result.Scheme)
	fmt.Printf("Type of scheme: %T\n", result.Scheme)
	fmt.Println("Host: ", result.Host)
	fmt.Printf("Type of scheme: %T\n", result.Scheme)
	fmt.Println("Path: ", result.Path)
	fmt.Printf("Type of scheme: %T\n", result.Scheme)
	fmt.Println("Port: ", result.Port())
	fmt.Printf("Type of scheme: %T\n", result.Scheme)
	fmt.Println("Raw Query: ", result.RawQuery)
	fmt.Printf("Type of scheme: %T\n", result.Scheme)
	qparams := result.Query()
	// query comes in the form of map[string][]string i.e maps formate key value pair
	fmt.Printf("Type of query params: %T\n", qparams)
	fmt.Println("Query Params: ", qparams)
	fmt.Println("Course Name: ", qparams["coursename"])
	fmt.Println("Payment ID: ", qparams["paymentid"])

	for key , val := range qparams{
		fmt.Printf("The key is %v and value is %v\n", key, val )
	}

	// always add & in the start to pass the actual value not the reference
	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "www.google.com",
		Path: "/tutcss",
		RawPath: "user=satvik",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println("Another URL is: ", anotherUrl)
}

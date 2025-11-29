package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://dummyjson.com/products"

func main() {
	fmt.Println("Welcome to web request using http")
	// get method in htttp

	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type %T\n", res)

	defer res.Body.Close() // closing the response body to prevent memory leaks

	data, err := io.ReadAll(res.Body) // io.ReadAll reads the data from the response body
	if err != nil {
		panic(err)
	}
	
	fmt.Println("Data is ", string(data))
}

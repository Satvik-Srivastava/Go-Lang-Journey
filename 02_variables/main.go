package main

import "fmt"

// jwtToken := "wfhiwfwf" // this format in go lang is not allowed, it is only allowed inside functions

var jwtToken string = "3rdqwdqd323"

const LoggingToken = "asdasdasdasd" // this variable is public because it starts with capital letter

func main() {
	// string value
	var username string = "satvik.com"
	fmt.Println("Hello, World from", username)
	fmt.Printf("Variable is of type: %T \n", username)

	// boolean value
	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	// small integer value
	var smallValue uint8 = 255 // uint8 can hold values from 0 to 255 only positive number 
	fmt.Println(smallValue)
	fmt.Printf("Variable is of type: %T \n", smallValue)

	// small negative integer value
	var smallNegativeValue int8 = -128
	fmt.Println(smallNegativeValue)
	fmt.Printf("Variable is of type: %T \n", smallNegativeValue)

	// float value
	var smallFloatValue float64 = 255.4554554
	fmt.Println(smallFloatValue)
	fmt.Printf("Variable is of type: %T \n", smallFloatValue)


	// default value and some alias
	var anotherValue int // default value is 0
	fmt.Println(anotherValue)
	fmt.Printf("Variable is of type: %T \n", anotherValue)

	var anotherstring string // default value is empty string
	fmt.Println(anotherstring)
	fmt.Printf("Variable is of type: %T \n", anotherstring)

	// implicit type
	var website = "satvik.com"
	fmt.Println(website)
	
	// now we will not use var

	websitename := "satvik.com"
	fmt.Println(websitename)
	websitename = "www.satvik.com"
	fmt.Println(websitename)

	fmt.Println("Logging Token is:", LoggingToken)
	fmt.Println("JWT Token is:", jwtToken)
}
package main

import "fmt"

func main() {
	fmt.Println("Hello, Structs!")
	//  There are no inheritance, super or parent child concept in Go

	satvik := User{"Satvik", 24, "satvik@gmail.com", true}
	fmt.Println(satvik) // {Satvik 24 satvik@gmail.com true}

	fmt.Printf("Satvik details are : %+v", satvik)  //{ %+v shows field names as well}
	// OUTPUT: Satvik details are : {Name:Satvik Age:24 Email:satvik@gmail.com Status:true}

	fmt.Printf("Name is %v", satvik.Name) //Name is Satvik
}

type User struct {
	// The first letter of the field name determines the visibility
	// If it starts with an uppercase letter, it is exported (public)
	Name   string
	Age    int
	Email  string
	Status bool
}

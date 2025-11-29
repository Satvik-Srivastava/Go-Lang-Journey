package main

import "fmt"

func main() {
	fmt.Println("Lecture on pointers")
	// There is a problem that occurs in most of the lang. when we store a value in a variable it is 
	// just a reference of the value. so when we pass this value from different function 
	// sometime it happens that we pass the copy of that value not the actual value which causes
	// irregularity in the output

	// so to avoid this irregualarity we use pointer. it give us the surity that the value you are passing
	// are not the value but the address of that value so any change in that value will be reflected	

	var ptr *int // declaration of pointer

	mynumber := 23
	
	ptr = &mynumber // assigning address of mynumber to pointer ptr, reference means &

	// & ke sath jo bhi lagta hai wo uska address deta hai

	// "*" ke sath jo bhi lagta hai wo uska value deta hai
	
	fmt.Println("Value of pointer ptr is: ", ptr) // by default pointer is nil
	fmt.Println("Value at address ptr is: ", *ptr) // dereferencing pointer means *

	*ptr = *ptr + 2 // changing the value at address ptr
	fmt.Println("New value of mynumber is: ", mynumber)
	fmt.Println("New value at address ptr is: ", &ptr)


}
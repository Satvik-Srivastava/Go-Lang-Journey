package main

import "fmt"

func main() {
	fmt.Println("Welcome in array")

	var fruits [4]string

	fruits[0] = "Apple"
	fruits[1] = "Mango"
	fruits[2] = "Banana"
	fruits[3] = "Grapes"	

	fmt.Println("Fruits list: ", fruits)
	fmt.Println("Length of fruits array is: ", len(fruits)) // if we have added only 3 values out of 4 it will still show the length as 4

	var vegies = [3]string{"Potato", "Brinjal", "Cauliflower"}

	fmt.Println("Vegies list: ", vegies)
	fmt.Println("Length of vegies array is: ", len(vegies))
}
package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices")

	// declaration of slice
	var fruitslice = []string{"orange", "banana", "apple"}

	fmt.Println("Fruits slice is: ", fruitslice)
	fmt.Printf("Type of fruitslice is : %T", fruitslice)

	// how to add values in slice (use append function)
	// variable = append(variable, value1, value2,...)

	fruitslice = append(fruitslice, "grapes", "kiwi")
	fmt.Println("\nFruits slice after append is: ", fruitslice)

	fruitslice = append(fruitslice[1:4]) // removing elements from slice
	fmt.Println("Fruits slice after removing first element: ", fruitslice)

	// use of make 
	highscores := make([]int ,4)
	highscores[0] = 2345
	highscores[1] = 4567
	highscores[2] = 6789
	highscores[3] = 7890

	// now if we try to add one more value it will give error because we have defined the length of slice as 4
	// highscores[4] = 8901 // this will give error

	// to avoid this we can use append function
	// now this append function will create a new slice and will reallocate all the values to new slice with new length
	highscores = append(highscores, 8901, 9999)
	fmt.Println("Highscores slice is: ", highscores)

	// sorting function in slice

	sort.Ints(highscores)
	fmt.Println("Highscores slice after sorting is: ", highscores)

	fmt.Println("Is highscores slice is sorted? ", sort.IntsAreSorted(highscores))

	// how to remove a value from slice based on index
	var courses = []string{"reactjs", "javascript", "html", "css", "java", "python"}
	fmt.Println("Courses slice before removing element: ", courses)

	index := 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("Courses slice after removing element: ", courses)


}
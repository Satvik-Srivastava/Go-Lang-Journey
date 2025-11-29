package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops")

	days := [] string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	// different ways to use for loops
	// for i:=0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for i := range days{
	// 	fmt.Printf("Day %v is %v\n", i+1, days[i])

	// }

	for index, day := range days{
		fmt.Printf("Day %v is %v\n", index+1, day)
	}

	// goto statement

	rogueValue := 1

	for rogueValue < 10 {
		if rogueValue == 5{
			rogueValue++
			goto lco
		}
		fmt.Println("Rogue Value is: ", rogueValue)
		rogueValue++
	}

	// declaring a label

	lco:
		fmt.Println("Jumped to lco")

}
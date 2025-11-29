package main

import (
	"fmt"	
)

func main() {
	fmt.Println("Welcome to maps")
	// declaration of map
	// variable := make(map[keydatatype]valuedatatype)
	languages := make(map[string] string)

	languages["js"] = "JavaScript"
	languages["py"] = "Python"
	languages["rb"] = "Ruby"

	fmt.Println("List of languages: ", languages)
	// particular value access
	fmt.Println("Value at key js is: ", languages["js"])

	languages["cpp"] = "C++" // adding new key value pair
	languages["pyt"] = "Python 3" // updating value at key py

	// delete value from map
	delete(languages, "rb")
	fmt.Println("List of languages after deleting rb: ", languages)

	// loop through map

	for key, value := range languages{
		fmt.Printf("Key is %s and value is %s \n", key, value)
	}

}
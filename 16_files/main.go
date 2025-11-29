package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Hello, understanding files in golang")
	content := "This is content from main.go file"

	file, err := os.Create("./mylocalgofile.txt") // creates a file in the current directory

	if err != nil {
		fmt.Println("There is some error")
		panic(err)
	}

	length, err := io.WriteString(file, content)
	if err != nil {
		fmt.Println("There is some error")
		panic(err)
	}

	fmt.Println("Length is: ", length)
	defer file.Close()

	readFile()

}


func readFile() {
	databyte, err := os.ReadFile("./mylocalgofile.txt")
	if err != nil {
		fmt.Println("There is some error")
		panic(err)
	}
	fmt.Printf("Type of the data %T\n", databyte)
	fmt.Println("Text data inside the file is: \n", string(databyte))
}

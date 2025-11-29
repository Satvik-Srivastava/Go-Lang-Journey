package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Hello, welcome to our restaurant!"
	reader := bufio.NewReader(os.Stdin)

	fmt.Println(welcome)
	fmt.Println("Please rate us now")

	// comma ok idiom
	input, err := reader.ReadString('\n')
	fmt.Println("Thanks for your rating, we got it as:", input)
	fmt.Printf("Type of this rating is %T\n", input)
	fmt.Printf("Type of this rating is %T", err)

}

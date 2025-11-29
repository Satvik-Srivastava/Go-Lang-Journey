package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our pizza app")
	fmt.Println("Please give us the rating:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for your rating, we got it as:", input)
	updaterating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	fmt.Println("Your updated rating is:", updaterating)
	fmt.Println("Your error:", err)


}

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to the time app")

	presenttime := time.Now()
	fmt.Println("The present time is:", presenttime)
	fmt.Println(presenttime.Format("01-08-2006"))

}

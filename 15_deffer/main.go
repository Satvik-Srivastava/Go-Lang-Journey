package main

import "fmt"

func main() {
	// the defer keyword make the thing after it to be executed at last just before the function ends
	fmt.Println("Hello")
	// when multiple defers are there they are executed in LIFO (last in first out) order
	defer fmt.Println("Deffered Hello 1")
	defer fmt.Println("Deffered Hello 2")
	defer fmt.Println("Deffered Hello 3") // order will be 3,2,1

	fmt.Println("World")

	mydeferFunction()
}

func mydeferFunction(){
	for i:=0; i<5; i++{
		defer fmt.Println(i) 
	}

}

// hello, world, 
// 0,1,2,3,4 => 4,3,2,1,0
// Deffered Hello 1,2,3 => deffered 3,2,1
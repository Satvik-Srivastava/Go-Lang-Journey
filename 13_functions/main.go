package main

import "fmt"

func fibo(a int, b int, size int){
	for i := 2; i < size; i++{
		nextIntem := a + b
		fmt.Printf("%v ", nextIntem)
		a = b
		b = nextIntem
	}
}

// if we dont know how many value will be passed we can use ... before datatype

func proAdder (values ...int) (int, string) {
	total := 0
	for _, num := range values{
		total += num
	}
	return total, "Sum calculated successfully"
}

func main() {
	fmt.Println("Welcome to functions")

	a, b := 0, 1

	fmt.Printf("%v %v ", a, b)
	fibo(a, b, 10)

	proTotal, msg := proAdder(1,2,3,4,5,6,7,8,9,10)
	fmt.Println("\nPro Total is: ", proTotal)
	fmt.Println("message is: ", msg)
}
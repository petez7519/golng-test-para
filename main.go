package main

import "fmt"

// Function to add two numbers
func add(a int, b int) int {
	return a + b
}

// Function to multiply two numbers
func multiply(a int, b int) int {
	return a * b
}

func main() {
	var num1, num2 int

	fmt.Print("Enter first number: ")
	fmt.Scan(&num1)

	fmt.Print("Enter second number: ")
	fmt.Scan(&num2)

	sum := add(num1, num2)
	product := multiply(num1, num2)

	fmt.Printf("The sum of %d and %d is %d\n", num1, num2, sum)
	fmt.Printf("The product of %d and %d is %d\n", num1, num2, product)
}

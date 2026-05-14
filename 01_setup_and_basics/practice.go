package main

import "fmt"

const PI float32 = 3.14

func hello() {
	fmt.Println("hola from español")
}

//Write a function add(a int, b int) int and call it from main.
func add(a, b int) int {
	return a + b
}

//Modify the function to return two values: sum and difference.
func arithmetic(x, y int) (int, int) {
	return x + y, x - y
}

//Write a function that swaps two strings
func swapStr(x, y string) (string, string) {
	return y, x
}

//convert int to float64
func convertType(a int) float64 {
	return float64(a)
}

//find type of variable using inferences
func findType(a int) {
	fmt.Printf("value of a is %d and it's type is %T\n", a, a)
}

//Write a function that returns whether a number is even or odd.
func isEven(a int) bool {
	return a%2 == 0
}

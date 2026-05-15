package main

import "fmt"

//Print numbers from 1 to 10 using a loop.
func printtens() {
	for i := 1; i < 11; i++ {
		fmt.Println(i)
	}
}

//Print all even numbers from 1 to 50.
func printevens() {
	for i := 1; i < 50; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Print(i)
	}
}

// Calculate the sum of first N natural numbers.
func sumofNNatural(N int) (sum int) {
	for N != 0 {
		sum += N
		N--
	}
	return
}

//Write a program to find factorial of a number.
func factorial(n int) int {
	fact := 1
	for i := 1; i <= n; i++ {
		fact *= i
	}
	return fact
}

//check if number is prime
func isPrime(a int) bool {
	for i := 2; i*i <= a; i++ {
		if a%i == 0 {
			return false
		}
	}
	return true
}

//demonstrate defer using statements like start, end, and differ
func showDefer() {
	defer fmt.Println("defer")
	fmt.Println("start")
	fmt.Println("end")
}

//program to find largest of the tree numbers
func findLargest(a, b, c int) int {
	max := a
	if max > b {
		max = b
	}
	if max > c {
		max = c
	}
	return max
}

func loopoverstring(str string) {
	for i, ch := range str {
		fmt.Println(i, ch)
	}
}

func practice() {
	fmt.Println("this is practice.go")
	printtens()
	// printevens()
	x := 5
	fmt.Println(sumofNNatural(x))
	fmt.Println(factorial(x))
	fmt.Println(isPrime(x))
	showDefer()
	loopoverstring("hari")
}

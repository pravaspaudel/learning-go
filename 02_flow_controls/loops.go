// loops,if,switch
package main

import (
	"fmt"
	"math"
	"runtime"
)

// go have only one looping control which is for
func findSum(n int) (sum int) {
	//standard loop
	// for i := n; i >= 0; i-- {
	// 	sum += i
	// }

	//infinite loops using for
	// i := n
	// for {
	// 	if i == 0 {
	// 		break
	// 	}
	// 	sum += i
	// 	i--
	// }

	//while loop
	i := n
	for i == 0 {
		sum += i
		i--
	}
	return
}

func findSqrt(x float64) string {
	if x < 0 {
		return fmt.Sprintf("%.6fi", math.Sqrt(-x))
	}
	return fmt.Sprintln(math.Sqrt(x))
}

// switch statements
func swtchstatments() {
	fmt.Print("your go os is ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("macOS")
	case "linux":
		fmt.Println("linux")
	default:
		fmt.Println(os)
	}
}

// defer statements
// defer statement delays the execution surrounding function returns
// these statements are pushed onto the stack after evaluating and called by popping from stack
func justdefer() {
	defer fmt.Println("this is defer1")
	defer fmt.Println("this is defer2")
	defer fmt.Println("this is defer3")
	defer fmt.Println("this is defer4")
	fmt.Println("this is not defer")
}

func loops() {
	fmt.Println("this is loops")
	n := 4
	fmt.Println(n, findSum(4))
	fmt.Println(findSqrt(5), findSqrt(-5))
	swtchstatments()
	justdefer()
}

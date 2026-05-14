package main

import "fmt"

func main() {
	hello()

	a, b := 4, 3
	fmt.Printf("addition of %d and %d is %v\n", a, b, add(a, b))

	add, sub := arithmetic(a, b)
	fmt.Printf("res is %d and %d\n", add, sub)

	str1, str2 := "hello", "world"
	fmt.Println(str1, str2)
	str1, str2 = swapStr(str1, str2)
	fmt.Println(str1, str2)

	//pi declared globally on practice.go
	fmt.Println("value of pi is", PI)

	x := 43
	fmt.Println(x, convertType(a)+3.2)

	findType(x)

	res := isEven(x)
	fmt.Println(x, res)

}

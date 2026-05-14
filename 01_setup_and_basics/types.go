package main

import "fmt"

//there are various types in go
//1. bool
//2. string
//3. int, int8, int16, int32, int64
//4. uint, uint8, uint16, uint32, uint64, uintptr
//5. float32,float64
//6. byte // alias for uint8
//7. complex64, complex128

// explict conversion of data types
func types() {
	var a int32 = 1024
	var i float32 = float32(a)
	const b int = 3
	fmt.Println(a, i+2.4, b)
}

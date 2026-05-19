package main

import "fmt"

//go doesnot have classes
//methods are the functions with special receiver arguments
//the receiver argument should be type (standard types like int,float not allowed)

type MyInt int

func (x MyInt) check() {
	if x > 0 {
		fmt.Println("x is greater than zero")
		return
	}
	fmt.Println("x is less than zero")
}

type Point struct {
	x float32
	y float32
}

//pointer receiver
func (p *Point) scale(sx, sy float32) {
	p.x = p.x * sx
}

func addtwo(a *int) int {
	return *a + 2
}

func methods() {
	fmt.Println("theser are the methods")

	var a MyInt
	a = 3
	a.check()

	pptr := &Point{3, 4}
	p := Point{3, 4}

	fmt.Println(p)
	p.scale(4, 5)
	fmt.Println(p)

	fmt.Println("pointers: ")
	fmt.Println(pptr)
	pptr.scale(4, 5)
	fmt.Println(pptr)
}

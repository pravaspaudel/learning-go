package main

import "fmt"

//pointers hold memory address of a value.
// unlike c, go has no pointer arithmetic

type Rect struct {
	l int
	b int
}

func pointers() {
	var p *int
	a := 54
	p = &a // pointer points to memory of a
	fmt.Println(p, *p)

	r := Rect{5, 9}
	//pointer to the struct
	var sptr *Rect
	sptr = &r
	fmt.Println(sptr, sptr.b)
}

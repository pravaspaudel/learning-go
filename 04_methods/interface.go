package main

import "fmt"

//interface is defined set of method signatures
//value of interface type can hold any value that implements those methods

type I interface {
	M()
}

type T struct {
	S string
}

//this means that this function implements interface's method
//type T implements M present in I
func (t *T) M() {
	fmt.Println(t.S)
}

func describe(i I) {
	fmt.Printf("(%v,%T)\n", i, i)
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v  is (%v years)\n", p.Name, p.Age)
}

func interfaces() {
	fmt.Println("interface.go")

	tptr := &T{S: "hari"}

	var a I
	a = tptr

	n, err := fmt.Print("hello\n")

	fmt.Println(n, err)

	describe(a)

	p1 := Person{"don'tknow", 30}
	p2 := Person{"didn'tknow", 35}

	fmt.Println(p1, p2)

}

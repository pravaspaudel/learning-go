package main

import "fmt"

// Create a Stack with methods: push , pop and peek
type Stack []int

func (s *Stack) push(val int) {
	*s = append(*s, val)
}

func (s *Stack) pop() (int, bool) {
	if len(*s) == 0 {
		return 0, false
	}
	idx := len(*s) - 1
	val := (*s)[idx]
	*s = (*s)[:idx]
	return val, true
}

func (s *Stack) peek() (int, bool) {
	if len(*s) == 0 {
		return 0, false
	}
	idx := len(*s) - 1
	return (*s)[idx], true
}

// Create a Queue with: Enqueue Dequeue
type queue []int32

func (q *queue) Enqueue(num int32) {
	*q = append(*q, num)
}

func (q *queue) Dequeue() (int32, bool) {
	if len(*q) == 0 {
		return 0, false
	}
	item := (*q)[0]
	*q = (*q)[1:]
	return item, true
}

//Create custom error type: type MyError struct { Msg string } Implement Error() method.
type CustomError struct {
	Msg string
}

func (err CustomError) Error() {
	fmt.Println("this is error", err.Msg)
}

//Define: type Celsius float64 and type Fahrenheit float64 Add two methods
// ToFahrenheit() and ToCelsius()
type Celsius float64
type Fahrenheit float64

func (f Fahrenheit) ToCelsius() Celsius {
	return Celsius(5 * (f - 32) / 9)
}
func (c Celsius) ToFahrenheit() Fahrenheit {
	return Fahrenheit(9*c/5 + 32)
}

//Create person struct and implement String() method (like fmt.Stringer)
type Personone struct {
	Name string
	age  int
}

func (p Personone) String() string {
	return fmt.Sprintf("person name is %v and age is %d", p.Name, p.age)
}

func practice0() {
	fmt.Println("practice0.go")
	p := Personone{"abc", 32}
	//just changed the way it prints the string using String()
	fmt.Println(p)

	var s Stack
	s.push(65)
	s.push(6)
	s.push(5)
	s.push(95)

	top, ok := s.peek()
	if ok {
		fmt.Println(top)
	} else {
		fmt.Println("stack is emtpy")
	}

	popped, ok := s.pop()
	if ok {
		fmt.Printf("%v was popped from stack", popped)
	}

}

package main

import (
	"fmt"
	"math"
	"slices"
	"strings"
)

// Create Processor interface: UpperCaseProcessor and ReverseProcessor
type Processor interface {
	Process(input string) string
}

type UpperCaseProcessor struct{}

func (u UpperCaseProcessor) Process(input string) string {
	return strings.ToUpper(input)
}

type ReverseProcessor struct{}

func (r ReverseProcessor) Process(input string) string {
	words := strings.Split(input, " ")
	slices.Reverse(words)
	return strings.Join(words, " ")
}

// Create PaymentMethod interface: CreditCard and PayPal
// Call Pay() polymorphically.
type PaymentMethod interface {
	Pay(amount float64)
}

type CreditCard struct{}
type Paypal struct{}

func (c CreditCard) Pay(amount float64) {
	fmt.Println("credit card used, paid", amount)
}

func (p Paypal) Pay(amount float64) {
	fmt.Println("paypal used amount = ", amount)
}

// Create Logger interface: with method   Log(message string)
// Implement: ConsoleLogger and  FileLogger (simulate with print)
type Logger interface {
	Log(message string)
}

type ConsoleLogger struct{}
type FileLogger struct{}

func (c ConsoleLogger) Log(message string) {
	fmt.Println("console :", message)
}

func (f FileLogger) Log(message string) {
	fmt.Println("file :", message)
}

// Create interface:Shape with Area() float64
// Implement for Rectangle, Circle.
type Shape interface {
	Area() float64
}

type Rectangle struct {
	Length  float64
	Breadth float64
}
type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Breadth
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// here we use s interface
func printArea(s Shape) {
	fmt.Println("the calculated area is", s.Area())
}

// middleware (function wrapper)
func Middleware(fn func()) func() {
	return func() {
		fmt.Println("start")
		fn()
		fmt.Println("end")
	}
}

func practice1() {
	fmt.Println("practice1.go")

	r := Rectangle{Length: 5.4, Breadth: 2.4}
	c := Circle{Radius: 3.4}

	fmt.Println(r.Area())
	fmt.Println(c.Area())

	shapes := []Shape{
		Rectangle{3, 3},
	}

	fmt.Println(shapes)

	rev1 := ReverseProcessor{}

	fmt.Println(rev1.Process("hi it's go"))

}

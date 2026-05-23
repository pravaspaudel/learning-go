package main

import (
	"fmt"
	"time"
)

// Launch a goroutine that prints numbers 1–10.
func printNumbers() {
	for i := 1; i < 11; i++ {
		fmt.Println(i)
	}
}

func practice0() {
	fmt.Println("practice0.go")

	go printNumbers()

	// Run two goroutines and print messages concurrently.
	go func() {
		for range 3 {
			fmt.Println("goroutine A")
			time.Sleep(time.Millisecond * 100)
		}
	}()

	go func() {
		for range 3 {
			fmt.Println("goroutine B")
			time.Sleep(time.Millisecond * 100)
		}
	}()

	// time.Sleep(1 * time.Second)

	// Use a channel to send a value from one goroutine to another.
	ch := make(chan int32) //unbuffered
	go func() {
		ch <- 43
	}()

	val := <-ch
	fmt.Println("this is go func with val", val)

	// Create a channel and send multiple values.
	ch2 := make(chan int)

	go func() {
		for i := 5; i <= 10; i++ {
			ch2 <- i
		}
		close(ch2)
	}()

	for v := range ch2 {
		fmt.Println(v)
	}

	// Use a buffered channel to store value from 10 to 14.
	ch3 := make(chan int, 5)
	ch3 <- 10
	ch3 <- 11
	ch3 <- 12

	fmt.Println(<-ch3, <-ch3, <-ch3)

	fmt.Println("finished")
}

package main

import (
	"fmt"
	"sync"
	"time"
)

// for concurrency, go has go subroutines . these are run by keyword go
// these go subroutines run concurrently with other functions
func say(s string) {
	fmt.Println("hello from go routine", s)
}

// channels are the way for communication between goroutines
// channel is a pipe through which goroutines can send and receive values
//channel := make(chan string)
// channel <- "hello"
// valueReceived := <-channel

//buffered channel vs unbuffered channel
//buffered channel has a capacity , it can hold values without receiver being ready yet
// ch:= make(chan int,3)

func sendData(ch chan string) {
	ch <- "this is a channel" //sending msg to channel
}

//we can set the direction of channel to only receive it or send it.
//chan string --> can send and receive (bidirectional)
//chan<-string --> can only send(send-only)
//<-chan string -->can only receive(receive-only)

func producer(ch chan<- string) {
	fmt.Println("sending data...")
	ch <- "data"
	fmt.Println("sent data")
}

func consumer(ch <-chan string) {
	fmt.Println("receiving data...")
	msg := <-ch
	fmt.Println("rece : ", msg)
}

//select statement let's goroutine wait on multiple channel operations at the same time
//and pick whichever is  ready first
// select{
// case msg1 := <-ch1:
// 	fmt.Println("received from ch1 ",msg1)

// case msg2 := <-ch2:
// 	fmt.Println("received from ch2",msg2)
// }
//whichever channel have data that one runs

// a waitgroup wait for a collection of goroutines to finish
//Add(x)  --> one more go routine starting
//Done()  --> one  goroutine finished
//Wait() --> wait until  the routine are finished
// waitgroup is like a counter which keeps count of the go routines and until counter is not zero it will wait

//sync.Mutex helps to prevent race condition
// we use mutex to lock and unlock the access value
//two functions Lock() and Unlock(). lock() will lock the variable and other routines have to wait for it

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Worker %d starting\n", id)
	// simulate work...
	fmt.Printf("Worker %d done\n", id)
}

func routines() {
	go say("h")
	time.Sleep(1 * time.Second)
	//need to be careful as if the main function exit before subroutine routines are killed
	fmt.Println("main function is done")

	// ch := make(chan string)
	// go sendData(ch)
	// //send blocks until someone receives it

	// msg := <-ch //main waits until something arrives
	// fmt.Println(msg)
	// fmt.Println("last ")

	ch := make(chan int)

	go func() {
		ch <- 30 //send in separate go routines
	}()

	time.Sleep(2000 * time.Millisecond)
	value := <-ch // it receives
	fmt.Println("function execution completed", value)

	ch2 := make(chan string)
	go producer(ch2)
	consumer(ch2)

	channel1 := make(chan string, 1)
	channel2 := make(chan string, 1)

	channel1 <- "one"
	channel2 <- "two"

	select {
	case msg := <-channel1:
		fmt.Println("Received from ch1:", msg)

	case msg := <-channel2:
		fmt.Println("received from ch2:", msg)
	}

	var wg sync.WaitGroup

	//we are throwing 3 go routines inside, on copeletion of each routine wg.Done() is executed
	for i := range 3 {
		wg.Add(i)
		go worker(i, &wg)
	}

	//wait unitls all 3 routines are executed
	wg.Wait()
	fmt.Println("all workers finished")

}

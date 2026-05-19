package main

import "fmt"

// Create an event system:
// Interface Listener
// Register multiple listeners
// Trigger event
// Create an event system where you define a listener concept.
// The system should allow multiple listeners to be registered, and
// when an event is triggered, all registered listeners should respond to it.

// //1. Event System

// Create an event system where you define a listener concept.
// The system should allow multiple listeners to be registered, and when an event is triggered, all registered listeners should respond to it.

// 2. Plugin System

// Design a plugin-based architecture using interfaces.
// The system should allow different plugin implementations to be swapped at runtime without changing the core logic.

// 3. Processing Pipeline

// Create a processing system where multiple processors can be chained together.
// Each processor should transform the input, and the output of one processor should become the input of the next.

// 4. Retry Mechanism

// Implement a retry system using functions.
// The function should attempt to execute an operation, and if it fails, it should retry a specified number of times.
// Simulate failures to test the retry behavior.

// 5. Mini Service System (Combined Task)

// Build a small service system that combines multiple concepts together.

// The system should include:

// Interfaces to define services and behaviors
// Structs to represent concrete implementations
// Methods to operate on those structures
// A map-based registry to store and retrieve services by name
// An event system to notify listeners when actions happen
// A plugin system where behavior can be swapped dynamically
// A processing pipeline to transform data step by step
// A retry mechanism to handle failures during execution

// The final system should simulate a real-world service architecture where all these components work together.

type Listener interface {
	Register() bool
	Trigger()
}

type User struct{}

// Build a plugin system:
// Interface-based
// Swap implementations at runtime
// Chain multiple processors (like pipeline processing).
// Implement retry logic using functions (simulate failures).
// Combine everything:
// Interface + struct + method + map
// Example: mini service system

func practice2() {
	fmt.Println("practice2.go")

}

package main

import "fmt"

// Higher-order function that takes a function as an argument
func apply(fn func(int) int, x int) int {
	return fn(x)
}

// A function to double a number
func double(x int) int {
	return x * 2
}

// Curried addition function
func curriedAdd(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

func main() {
	// First Class and Higher Order Functions
	result := apply(double, 5)
	fmt.Println("First Class and Higher Order Functions:")
	fmt.Printf("Result of applying double: %d\n", result)
	// Output:
	// First Class and Higher Order Functions:
	// Result of applying double: 10

	fmt.Println()

	// Currying
	addFive := curriedAdd(5)
	curriedResult := addFive(3)
	fmt.Println("Currying:")
	fmt.Printf("Result of curried addition: %d\n", curriedResult)
	// Output:
	// Currying:
	// Result of curried addition: 8

	fmt.Println()

	// Closures
	c := counter()
	fmt.Println("Closures:")
	fmt.Printf("Counter value: %d\n", c())
	fmt.Printf("Counter value: %d\n", c())
	// Output:
	// Closures:
	// Counter value: 1
	// Counter value: 2

	fmt.Println()

	// Defer
	fmt.Println("Defer:")
	fmt.Println("Regular statement")
	defer fmt.Println("Deferred statement executed")
	// Output:
	// Defer:
	// Regular statement
	// Deferred statement executed

	fmt.Println()

	// Anonymous Functions
	add := func(a, b int) int {
		return a + b
	}
	anonymousResult := add(3, 4)
	fmt.Println("Anonymous Functions:")
	fmt.Printf("Result of anonymous addition: %d\n", anonymousResult)
	// Output:
	// Anonymous Functions:
	// Result of anonymous addition: 7
}

// Counter function closure
func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

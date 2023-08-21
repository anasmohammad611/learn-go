package main

import "fmt"

func main() {
	var input int

	fmt.Print("Enter something: ")
	_, err := fmt.Scan(&input)
	if err != nil {
		return
	}

	fizzBuzz(input)
}

func fizzBuzz(inp int) {
	for i := 1; i <= inp; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Print("fizzBuzz")
		} else if i%3 == 0 {
			fmt.Print("fizz")
		} else if i%5 == 0 {
			fmt.Print("Buzz")
		} else {
			fmt.Print(i)
		}
		fmt.Print(" ")
	}
	fmt.Println()
}

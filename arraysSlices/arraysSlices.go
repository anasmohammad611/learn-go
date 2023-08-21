package main

import (
	"fmt"
)

func main() {
	// Arrays
	fmt.Println("Arrays:")
	var arr [5]int // Declare an array of 5 integers
	arr[0] = 10    // Assign values to elements
	arr[1] = 20
	fmt.Println("Array:", arr) // Print the array

	// Slices
	fmt.Println("Slices:")
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Original Slice:", slice)

	// Append to a slice
	slice = append(slice, 6, 7, 8)
	fmt.Println("After Append:", slice)

	// Copy a slice
	newSlice := make([]int, len(slice))
	copy(newSlice, slice)
	fmt.Println("Copied Slice:", newSlice)

	// Iterate over a slice using range
	fmt.Println("Iterating using range:")
	for index, value := range slice {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// Slices using make
	fmt.Println("\nSlices using make:")
	sliceMake := make([]int, 3, 5)
	sliceMake[0] = 100
	sliceMake[1] = 200
	sliceMake[2] = 300
	fmt.Println("Slice using make:", sliceMake)

	// Len and Cap of slice
	fmt.Println("\nLen and Cap:")
	fmt.Println("Length of sliceMake:", len(sliceMake))
	fmt.Println("Capacity of sliceMake:", cap(sliceMake))

	// Variadic Function
	fmt.Println("\nVariadic Function:")
	numbers := []int{1, 2, 3, 4, 5}
	sum := calculateSum(numbers...)
	fmt.Println("Sum of numbers:", sum)
}

// Variadic function to calculate the sum of integers
func calculateSum(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

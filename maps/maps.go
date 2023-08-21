package main

import (
	"fmt"
)

func main() {
	// Create an empty map
	myMap := make(map[string]int)

	// Add key-value pairs
	myMap["one"] = 1
	myMap["two"] = 2
	myMap["three"] = 3

	// Print the map
	fmt.Println("Initial Map:", myMap) // Output: Initial Map: map[one:1 three:3 two:2]

	// Access a value using key
	value := myMap["two"]
	fmt.Println("Value for key 'two':", value) // Output: Value for key 'two': 2

	// Check if a key exists
	_, exists := myMap["four"]
	fmt.Println("'four' exists:", exists) // Output: 'four' exists: false

	// Delete a key-value pair
	delete(myMap, "three")
	fmt.Println("Map after deleting 'three':", myMap) // Output: Map after deleting 'three': map[one:1 two:2]

	// Modify an existing value
	myMap["two"] = 22
	fmt.Println("Map after modifying 'two':", myMap) // Output: Map after modifying 'two': map[one:1 two:22]

	// Iterating over the map
	fmt.Println("Iterating over the map:")
	for key, value := range myMap {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}
	// Output:
	// Iterating over the map:
	// Key: one, Value: 1
	// Key: two, Value: 22

	// Using different key types
	keyTypes := make(map[interface{}]string)
	keyTypes[1] = "integer"
	keyTypes["two"] = "string"
	fmt.Println("Map with different key types:", keyTypes)
	// Output: Map with different key types: map[1:integer two:string]

	// Map mutations
	map1 := make(map[string]int)
	map1["a"] = 1

	map2 := map1
	map2["b"] = 2

	fmt.Println("Map1:", map1) // Output: Map1: map[a:1 b:2]
	fmt.Println("Map2:", map2) // Output: Map2: map[a:1 b:2]

	// Modify a value in map1
	map1["a"] = 3
	fmt.Println("Map1 after modifying 'a':", map1)
	// Output: Map1 after modifying 'a': map[a:3 b:2]
	fmt.Println("Map2 after modifying 'a' in map1:", map2)
	// Output: Map2 after modifying 'a' in map1: map[a:3 b:2]
}

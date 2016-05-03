package main

import "fmt"

// Type can be supplied after each parameter...
func add(x int, y int) int {
	return x + y
}
// ...Or if both parameters are the same type, at the end.
func sum(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(123,234))
	fmt.Println(sum(234,123))
}

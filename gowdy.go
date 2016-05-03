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

// A function can return any number of results, which seems super
// cool. I know I've had to create arrays before just to get two
// pieces of data back. I'm pretty sure multiline comments work now.
func addsum(a, b int) (int, int, int) {
	var x int
	x = a + b
	return a, b, x
}

func main() {
	a, b, c := addsum(add(123,234), 124)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(add(123,234))
	fmt.Println(sum(234,123))
}

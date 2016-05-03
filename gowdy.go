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

func sumadd(a, b int) (x, y int) {
	x = a + b
	y = b - a
	return
}

func main() {
	a, b, c := addsum(add(123,234), 124)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// If you declare something like "second", you have to use it otherwise
	// a build error will fire.
	first, second := sumadd(321,320)
	fmt.Println("First: ", first)
	fmt.Println("Second: ", second)

	// If I print sumadd directly, both return values output
	fmt.Println(sumadd(321,320))

	// However, something like this won't work because it treats this as
	// multi-value in single-value context. ?
	//fmt.Println("First and Second: ", sumadd(321,320))

	// This does work though because a single value is returned.
	fmt.Println("Add: ", add(123,234))

	fmt.Println(sum(234,123))
}

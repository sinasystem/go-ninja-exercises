package main

import "fmt"

var x int
var y string
var z bool

func main() {
	x, y, z = 42, "James Bond", true

	s := fmt.Sprintf("%v,%v,%v", x, y, z)

	fmt.Println(s)
}

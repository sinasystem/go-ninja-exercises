package main

import "fmt"

type age int

var x age
var y int

func main() {
	fmt.Printf("x: %v - %T\n", x, x)
	x = 42
	fmt.Printf("x: %v", x)

	y = int(x)
	fmt.Printf("y: %v - %T\n", y, y)

}

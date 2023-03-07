package main

import "fmt"

type age int

var x age

func main() {
	fmt.Printf("x: %v - %T\n", x, x)
	x = 42
	fmt.Println(x)

}

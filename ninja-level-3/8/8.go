package main

import "fmt"

func main() {
	x := 1111
	switch {
	case x < 100:
		fmt.Println("Less than 100")
	case x == 100:
		fmt.Println("100")
	default:
		fmt.Println("More than 100")

	}
}

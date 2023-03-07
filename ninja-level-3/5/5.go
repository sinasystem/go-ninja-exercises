package main

import "fmt"

func main() {
	for i := 11; i < 100; i++ {
		fmt.Printf("%v mod 4 = %v\n", i, i%4)
	}
}

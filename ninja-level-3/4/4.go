package main

import "fmt"

func main() {
	x := 1997

	for {
		if x > 2023 {
			break
		}
		fmt.Println(x)
		x++
	}
}

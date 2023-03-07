package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sumXi := foo(xi...)
	fmt.Println(sumXi)
}
func foo(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

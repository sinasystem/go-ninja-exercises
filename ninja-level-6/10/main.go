package main

import (
	"fmt"
	"math"
)

func main() {

	fourTimes2 := multiplyby2(3)
	fmt.Println(fourTimes2())
	fmt.Println(fourTimes2())
	fmt.Println(fourTimes2())
	fmt.Println(fourTimes2())
	fmt.Println(fourTimes2())
	fmt.Println(fourTimes2())
	fmt.Println(fourTimes2())

}
func multiplyby2(inp int) func() int {
	x := 0
	return func() int {
		fmt.Printf("%d * 2^%d\n", inp, x)
		out := inp * int(math.Pow(2, float64(x)))
		x++
		return out
	}
}

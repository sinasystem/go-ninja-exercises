package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	sli1 := x[:5]
	sli2 := x[5:]
	sli3 := x[2:7]
	sli4 := x[1:6]
	fmt.Println(sli1)
	fmt.Println(sli2)
	fmt.Println(sli3)
	fmt.Println(sli4)

}

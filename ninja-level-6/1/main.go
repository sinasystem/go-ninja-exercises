package main

import "fmt"

func main() {
	i1 := foo()
	i2, str := bar()

	fmt.Println(i1)
	fmt.Println(i2, str)

}
func foo() int {
	return 85
}
func bar() (int, string) {
	return 69, "Joon"
}

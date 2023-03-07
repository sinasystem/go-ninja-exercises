package main

import "fmt"

func main() {
	fmt.Println("First This")
	foo()
	defer bruh()
	fmt.Println("The Last one")
}
func foo() {
	defer bar()
	fmt.Println("Hi from 'foo'")

}
func bar() {
	fmt.Println("Hi from 'bar'")
}
func bruh() {
	fmt.Println("No I'm the last one!")
}

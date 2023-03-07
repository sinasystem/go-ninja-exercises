package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("Hi! I'm", p.first, p.last, ". I'm", p.age, "years old.")
}
func main() {
	p1 := person{
		first: "Sina",
		last:  "Nouri",
		age:   24,
	}
	p1.speak()
}

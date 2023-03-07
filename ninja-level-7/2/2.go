package main

import "fmt"

type person struct {
	name, address string
	age           int
}

func changeme(p *person, addr string) {
	(*p).address = addr
}
func main() {
	p1 := person{
		name:    "Sina",
		address: "Lahijan",
		age:     24,
	}
	fmt.Printf("p1: %v\n", p1)
	changeme(&p1, "Tehran")
	fmt.Printf("new p1: %v\n", p1)

}

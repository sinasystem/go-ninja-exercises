package main

import "fmt"

type Person struct {
	Name string
	Age  int
	Sex  string
}
type Human interface {
	Speak()
}

func (p *Person) Speak() {
	fmt.Printf("%s says `Blah Blah Blah`", p.Name)
}

func saySomething(h Human) {
	h.Speak()
}
func main() {
	p1 := Person{
		Name: "Sina Nouri",
		Age:  24,
		Sex:  "M",
	}
	saySomething(&p1)
	// saySomething(p1)		doesn't work

}

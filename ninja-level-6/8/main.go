package main

import "fmt"

func main() {
	x := greetMaker()

	fmt.Println(x("Sina"))
}
func greetMaker() func(name string) string {
	return func(name string) string {
		return fmt.Sprintf("Greetings from %s", name)
	}
}

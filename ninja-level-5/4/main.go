package main

import "fmt"

func main() {

	sina := struct {
		first string
		last  string
		age   int
	}{
		first: "Sina",
		last:  "noori",
		age:   24,
	}
	fmt.Println(sina)
}

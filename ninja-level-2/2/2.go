package main

import "fmt"

func main() {
	g := 42 == 44
	h := 42 <= 44
	i := 42 >= 44
	j := 42 != 44
	k := 42 < 44
	l := 42 > 44
	fmt.Print(g, h, i, j, k, l)
}

package main

import (
	"fmt"
)

const (
	_               = iota
	nextYear        = iota + 2023
	twoYearsLater   = iota + 2023
	threeYearsLater = iota + 2023
	fourYearsLater  = iota + 2023
)

func main() {
	fmt.Println(nextYear, twoYearsLater, threeYearsLater, fourYearsLater)
}

package main

import "fmt"

func main() {
	mynumber := 85
	fmt.Printf("Dec %d, Bin %b, Hex %#x\n", mynumber, mynumber, mynumber)
	shifted := mynumber << 1
	fmt.Printf("Dec %d, Bin %b, Hex %#x\n", shifted, shifted, shifted)

}

package main

import "fmt"

func main() {
	favSport := "e-sports"
	switch favSport {
	case "chess":
		fmt.Println("Boring")
	case "swimming":
		fmt.Println("Gay")
	case "e-sports":
		fmt.Println("Weeb")
	case "football":
		fmt.Println("Braindead")
	}
}

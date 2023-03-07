package main

import "fmt"

func main() {
	x := [][]string{
		{"James", "Bond", "Shaken, not stirred"},
		{"Miss", "Moneypenny", "Helloooooo, James."},
	}
	for i, v := range x {
		fmt.Printf("Record %d\n", i)
		for j := range v {
			fmt.Println(v[j])
		}
		fmt.Println()
	}

	fmt.Println("------------------------")
	fmt.Println("Another Way")
	fmt.Println("------------------------")

	z1 := []string{"James", "Bond", "Shaken, not stirred"}
	z2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}

	zzz := [][]string{z1, z2}

	for i, v := range zzz {
		fmt.Printf("Record %d\n", i)
		for j := range v {
			fmt.Println(v[j])
		}
		fmt.Println()
	}

}

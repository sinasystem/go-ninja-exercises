package main

import "fmt"

func main() {
	x := map[string][]string{
		`bond_james`:      {`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being evil`, `Ice cream`, `Sunsets`},
	}
	for k, v := range x {
		fmt.Println(k)
		for i := range v {
			fmt.Printf("\t\t%d\t%s\n", i, v[i])
		}
	}
}

package main

import "fmt"

func main() {
	x := map[string][]string{
		`bond_james`:      {`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being evil`, `Ice cream`, `Sunsets`},
	}

	x[`nourossana_sina`] = []string{`game`, `sex`, `youtube`}

	delete(x, "no_dr")
	for k, v := range x {
		fmt.Printf("%s\t%v\n", k, v)
	}

}

package main

import "fmt"

type person struct {
	first, last  string
	favIceCreams []string
}

func main() {
	p1 := person{
		first:        "Sina",
		last:         "Noori",
		favIceCreams: []string{"vanilla", "chocolate", "dark chocolate"},
	}
	p2 := person{
		first:        "Mahtab",
		last:         "Foo",
		favIceCreams: []string{"double chocolate", "cherry"},
	}
	fmt.Println(p1)
	for i := range p1.favIceCreams {
		fmt.Println(p1.favIceCreams[i])
	}
	fmt.Println(p2)

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}
	for _, v := range m {
		fmt.Println(v.first, v.last)
		for i := range v.favIceCreams {
			fmt.Println("\t", v.favIceCreams[i])
		}

	}
}

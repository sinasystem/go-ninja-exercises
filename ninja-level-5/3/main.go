package main

import "fmt"

type vehicle struct {
	doors int
	color string
}
type truck struct {
	vehicle
	fourWheel bool
}
type sedan struct {
	vehicle
	luxury bool
}

func main() {
	tr1 := truck{
		vehicle: vehicle{
			doors: 2,
			color: "grey",
		},
		fourWheel: true,
	}
	sed1 := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "orange",
		},
		luxury: true,
	}
	fmt.Println(tr1)
	fmt.Println(sed1)
	fmt.Println(tr1.doors)
	fmt.Println(sed1.color)

}

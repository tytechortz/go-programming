package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type sedan struct {
	vehicle
	luxury bool
}

type truck struct {
	vehicle
	fourWheel bool
}

func main() {
	s := sedan {
				vehicle: vehicle{
					doors: 4,
					color: "black",
				},
				luxury: true,
	}
		

	t := truck {
				vehicle: vehicle{
					doors: 2,
					color: "black",
				},
				fourWheel: true,
	}
		
	fmt.Println(s)
	fmt.Println(t)
	fmt.Println(t.doors)
	fmt.Println(s.doors)

}
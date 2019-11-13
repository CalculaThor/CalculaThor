package main

import (
	"fmt"
)

func main() {
	SetPoints([]Point{
		Point{1, 0.6747},
		Point{1.2, 0.8491},
		Point{1.4, 1.1214},
		Point{1.6, 1.4921},
		Point{1.8, 1.9607},
		Point{2, 2.5258},
	})
	fmt.Println(NewtonDifDiv())
	lagrange([]Point{
		Point{1,2},
		Point{2,3},
		Point{3,4},
	})
}
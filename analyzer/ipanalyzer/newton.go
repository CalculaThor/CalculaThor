package main 

type Point struct {
	X, Y float64
}

var points []Point

func NewtonDifDiv() float64 {
	return newtonDifDivAux(points)
}

func newtonDifDivAux(set []Point) (ans float64) {
	n := len(set)
	if n == 1 {
		ans = set[0].Y
		return 
	}
	v1 := newtonDifDivAux(set[0:n - 1])
	v2 := newtonDifDivAux(set[1:n])

	ans = (v2-v1)/(set[n-1].X-set[0].X)
	return 
}

func SetPoints(p []Point) {
	points = p
}
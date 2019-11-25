package ipanalyzer

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

var points []Point
var table, nevilleTable [][]float64
var okTable, okNevTable [][]bool
var badIn bool

func SetPoints(p []Point) {
	badIn = false
	points = p
	n := len(p)
	table = make([][]float64, n)
	nevilleTable = make([][]float64, n)
	okTable = make([][]bool, n)
	okNevTable = make([][]bool, n)
	for i := 0; i < n; i++ {
		table[i] = make([]float64, n)
		nevilleTable[i] = make([]float64, n)
		okTable[i] = make([]bool, n)
		okNevTable[i] = make([]bool, n)
	}
}

func InterpolateNewton(x float64) float64 {
	n := len(points)
	ans, mul := table[0][0], float64(1)
	for i := 1; i < n; i++ {
		mul *= x - points[i-1].X
		ans += table[i][i] * mul
	}
	return ans

}

func NewtonDifDiv() {
	newtonDifDivAux(points, 0, len(points)-1)
}

func newtonDifDivAux(set []Point, a, b int) (ans float64) {
	n := len(set)
	if n == 1 {
		if okTable[b][0] {
			ans = table[b][0]
		} else {
			ans = set[0].Y
			table[b][0] = ans
			okTable[b][0] = true
		}
		return
	}
	var v1, v2 float64
	if okTable[b-1][n-2] {
		v1 = table[b-1][n-2]
	} else {
		v1 = newtonDifDivAux(set[0:n-1], a, b-1)
	}
	if okTable[b][n-2] {
		v2 = table[b][n-2]
	} else {
		v2 = newtonDifDivAux(set[1:n], a+1, b)
	}
	if set[n-1].X == set[0].X {
		badIn = true
	}
	ans = (v2 - v1) / (set[n-1].X - set[0].X)
	table[b][n-1] = ans
	okTable[b][n-1] = true
	return
}

func LagrangeString() (ans string) {
	n := len(points)
	var num string
	den := float64(1)
	for i := 0; i < n; i++ {
		sign := "+"
		if points[i].Y < 0 {
			sign = "-"
		} else if i == 0 {
			sign = ""
		}
		num += fmt.Sprintf("%s(%g", sign, math.Abs(points[i].Y))
		for j := 0; j < n; j++ {
			if i != j {
				num += fmt.Sprintf("(x%+g)", -points[j].X)
				den *= (points[i].X - points[j].X)
			}
		}
		ans += fmt.Sprintf("%s/%g)", num, den)
		if den == 0 {
			badIn = true
		}
		num = ""
		den = 1
	}
	return
}

func LagrangeInterpolation(x float64) (ans float64) {
	n := len(points)
	for k := 0; k < n; k++ {
		mul := float64(1)
		for i := 0; i < n; i++ {
			if i != k {
				if points[k].X == points[i].X {
					badIn = true
				}
				mul *= (x - points[i].X) / (points[k].X - points[i].X)
			}
		}
		ans += mul * points[k].Y
	}
	return
}

func NevilleInterpolation(x float64) float64 {
	return nevilleInterpolationAux(x, points[:], 0, len(points)-1)
}

func nevilleInterpolationAux(x float64, set []Point, a, b int) (ans float64) {
	n := len(set)
	if n == 1 {
		if okNevTable[b][0] {
			ans = nevilleTable[b][0]
		} else {
			ans = set[0].Y
			nevilleTable[b][0] = ans
			okNevTable[b][0] = true
		}
		return
	}
	var v1, v2 float64
	if okNevTable[b-1][n-2] {
		v1 = nevilleTable[b-1][n-2]
	} else {
		v1 = nevilleInterpolationAux(x, set[0:n-1], a, b-1)
	}
	if okNevTable[b][n-2] {
		v2 = nevilleTable[b][n-2]
	} else {
		v2 = nevilleInterpolationAux(x, set[1:n], a+1, b)
	}
	if set[n-1].X == set[0].X {
		badIn = true
	}
	ans = (v2*(x-set[0].X) - v1*(x-set[n-1].X)) / (set[n-1].X - set[0].X)
	nevilleTable[b][n-1] = ans
	okNevTable[b][n-1] = true
	return
}

func GetNewtonTable() [][]float64 {
	return table
}

func GetNevilleTable() [][]float64 {
	return nevilleTable
}

func BadIn() bool {
	return badIn
}

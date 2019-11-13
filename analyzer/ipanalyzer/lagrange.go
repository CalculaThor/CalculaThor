package main
import (
	"fmt"
)
func lagrange(set []Point){
	n := len(set)
	var numerador string
	denominador := float64(1)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i!=j {
				numerador+= fmt.Sprintf("(x-%f)", set[j].X)
				denominador*= (set[i].X-set[j].X)
			}
		}
		fmt.Printf("%s/%g\n",numerador,denominador)
		numerador= ""
		denominador=1
	}
}
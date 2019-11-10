package analyzer

import (
	"github.com/CalculaThor/CalculaThor/analyzer/svanalyzer"
	"github.com/dtylman/gowd"
	"math"
	"strconv"
)

var em map[string]*gowd.Element

func BeginAnalyzer(mp map[string]*gowd.Element) {
	em = mp
	em["submit_button"].OnEvent("onclick", initSolution)
}

func initSolution(sender *gowd.Element, event *gowd.EventElement) {
	switch em["type_selector"].GetValue() {
	case "single":
		solveSingleEquation()
	case "system":

	case "inter":

	default:
	}
}

func solveSingleEquation() {
	f := em["single_problem_def"].GetValue()
	_ := svanalyzer.SetF(f)
	switch em["single_method_selector"].GetValue() {
	case "bisection":
		a, _ := strconv.ParseFloat(em["x1"].GetValue(), 64)
		b, _ := strconv.ParseFloat(em["x2"].GetValue(), 64)
		tol, _ := strconv.ParseFloat(em["tol"].GetValue(), 64)
		maxit, _ := strconv.ParseInt(em["it"].GetValue(), 10, 32)
		_ := svanalyzer.Bisection(a, b, tol, uint(math.Abs(float64(maxit))))


	case "false_pos":
		a, _ := strconv.ParseFloat(em["x1"].GetValue(), 64)
		b, _ := strconv.ParseFloat(em["x2"].GetValue(), 64)
		tol, _ := strconv.ParseFloat(em["tol"].GetValue(), 64)
		maxit, _ := strconv.ParseInt(em["it"].GetValue(), 10, 32)
		_ := svanalyzer.FalsePosition(a, b, tol, uint(math.Abs(float64(maxit))))
	case "secant":

	case "search":

	case "fixed_point":

	case "newton":

	case "multi":

	}
}

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
	err := svanalyzer.SetF(f)
	if err != nil {
		return
	}
	panel := em["solution_panel"]
	panel.RemoveElements()
	panel.Show()
	switch em["single_method_selector"].GetValue() {
	case "bisection":

		a, _ := strconv.ParseFloat(em["x1"].GetValue(), 64)
		b, _ := strconv.ParseFloat(em["x2"].GetValue(), 64)
		tol, _ := strconv.ParseFloat(em["tol"].GetValue(), 64)
		maxit, _ := strconv.ParseInt(em["it"].GetValue(), 10, 32)
		ans := svanalyzer.Bisection(a, b, tol, uint(math.Abs(float64(maxit))))
		if !ans.BadIn && !ans.IsAlmostRoot && !ans.IsRoot {
			panel.AddElement(gowd.NewElement("h2")).SetValue("Not enough iterations.")
		} else if ans.IsRoot {
			panel.AddElement(gowd.NewElement("h2")).SetValue(strconv.FormatFloat(ans.Root, 'E', 10, 64))
		} else if ans.IsAlmostRoot {
			panel.AddElement(gowd.NewElement("h2")).SetValue(strconv.FormatFloat(ans.Root, 'E', 10, 64))
		} else if ans.BadIn {
			panel.AddElement(gowd.NewElement("h2")).SetValue("Bad In.")
		}

	case "false_pos":
		a, _ := strconv.ParseFloat(em["x1"].GetValue(), 64)
		b, _ := strconv.ParseFloat(em["x2"].GetValue(), 64)
		tol, _ := strconv.ParseFloat(em["tol"].GetValue(), 64)
		maxit, _ := strconv.ParseInt(em["it"].GetValue(), 10, 32)
		svanalyzer.FalsePosition(a, b, tol, uint(math.Abs(float64(maxit))))
	case "secant":

	case "search":

	case "fixed_point":

	case "newton":

	case "multi":

	}
}

package analyzer

import (
	"github.com/dtylman/gowd"
)

var em map[string]*gowd.Element

func BeginAnalyzer(mp map[string]*gowd.Element) {
	em = mp
	em["submit_button"].OnEvent("onclick", initSolution)
	em["abs"].OnEvent("onclick", setAbs)
	em["rel"].OnEvent("onclick", setRel)
	em["SE_rel"].OnEvent("onclick", setRelNorm)
	em["SE_abs"].OnEvent("onclick", setAbsNorm)
	em["infinity_norm"].OnEvent("onclick", setInfinity)
	em["euclidean_norm"].OnEvent("onclick", setEuclidean)
}

func initSolution(sender *gowd.Element, event *gowd.EventElement) {
	em["solution_panel"].Show()
	switch em["type_selector"].GetValue() {
	case "single":
		solveSingleEquation()
	case "system":
		solveSystemOfEquations()
	case "inter":

	default:
	}
}

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
	em["SE_rel"].OnEvent("onclick", setRel)
	em["SE_abs"].OnEvent("onclick", setRel)
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

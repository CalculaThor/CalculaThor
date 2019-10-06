package gui

import "github.com/dtylman/gowd"

type singleVarSolutionPanel struct {
	*gowd.Element
	table  *gowd.Element
	result *gowd.Element
	plt    *gowd.Element
}

func solve() *singleVarSolutionPanel {
	p := &singleVarSolutionPanel{}
	p.Element = em["solution_panel"]
	p.plt = em["plot"]
	p.table = em["result_table"]
	p.result = em["result"]

	return p
}

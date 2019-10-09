package gui

import (
	"github.com/dtylman/gowd"
)

type singleSolutionPanel struct {
	*gowd.Element
	table  *gowd.Element
	result *gowd.Element
	plt    *gowd.Element
}

func (p *singleSolutionPanel) beginSolution() {
	p.Element = em["solution_panel"]
	p.plt = em["plot"]
	p.table = em["result_table"]
	p.result = em["result"]
	p.Hide()

}

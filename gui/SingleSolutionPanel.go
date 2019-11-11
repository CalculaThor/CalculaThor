package gui

import (
	"github.com/dtylman/gowd"
)

type SingleSolutionPanel struct {
	*gowd.Element
}

func (p *SingleSolutionPanel) beginSolution() {
	p.Element = em["solution_panel"]
	p.Hide()

}

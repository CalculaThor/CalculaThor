package gui

import (
	"github.com/dtylman/gowd"
)

type problemPanel struct {
	*gowd.Element
	definition *gowd.Element
	problem    *gowd.Element
	method     *gowd.Element
	in         *gowd.Element
}

func beginPanel() *problemPanel {
	p := &problemPanel{}
	p.Element = em["problem_panel"]
	p.definition = em["what"]
	p.method = em["how"]
	p.problem = em["which"]
	p.in = em["data"]

	p.method.Hide()
	p.problem.Hide()
	p.in.Hide()

	em["gin"].Hide()
	em["dfin"].Hide()
	em["d2fin"].Hide()

	em["type_selector"].OnEvent("onchange", p.openProblem)
	em["single_problem_def"].OnEvent("onchange", p.openMethods)
	em["single_method_selector"].OnEvent("onchange", p.openAdditionalInfo)

	return p
}

func (p *problemPanel) openProblem(sender *gowd.Element, event *gowd.EventElement) {
	switch sender.GetValue() {
	case "single":
		p.problem.Show()
	}
}

func (p *problemPanel) openMethods(sender *gowd.Element, event *gowd.EventElement) {
	switch em["type_selector"].GetValue() {
	case "single":
		//TODO: equation verification
		if p.method.Hidden {
			p.method.Show()
		}
	}
}

func (p *problemPanel) openAdditionalInfo(sender *gowd.Element, event *gowd.EventElement) {
	switch sender.GetValue() {
	case "none", "bisection", "false_pos", "secant", "search":
		em["gin"].Hide()
		em["dfin"].Hide()
		em["d2fin"].Hide()
		p.in.Show()
	case "fixed_point":
		em["gin"].Show()
		em["dfin"].Hide()
		em["d2fin"].Hide()
		p.in.Hide()
	case "newton":
		em["gin"].Hide()
		em["dfin"].Show()
		em["d2fin"].Hide()
		p.in.Hide()
	case "multi":
		em["gin"].Hide()
		em["dfin"].Show()
		em["d2fin"].Show()
		p.in.Hide()

	}
}

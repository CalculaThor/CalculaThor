package gui

import (
	"github.com/dtylman/gowd"
)

type SingleProblemPanel struct {
	*gowd.Element
	definition *gowd.Element
	problem    *gowd.Element
	method     *gowd.Element
	in         *gowd.Element
}

func (p *SingleProblemPanel) beginProblem() {
	p.Element = em["problem_panel"]
	p.definition = em["what"]
	p.method = em["how"]
	p.problem = em["which"]
	p.in = em["single_data"]

	p.method.Hide()
	p.problem.Hide()
	p.in.Hide()

	em["gin"].Hide()
	em["dfin"].Hide()
	em["d2fin"].Hide()

	em["type_selector"].OnEvent("onchange", p.openProblem)
	em["single_problem_def"].OnEvent("onchange", p.openMethods)
	em["single_method_selector"].OnEvent("onchange", p.openAdditionalInfo)
	em["g"].OnEvent("onchange", p.gfuncentry)
	em["df"].OnEvent("onchange", p.dfentry)
	em["d2f"].OnEvent("onchange", p.d2fentry)

}

func (p *SingleProblemPanel) openProblem(sender *gowd.Element, event *gowd.EventElement) {
	switch sender.GetValue() {
	case "single":
		p.problem.Show()
	}
}

func (p *SingleProblemPanel) openMethods(sender *gowd.Element, event *gowd.EventElement) {
	switch em["type_selector"].GetValue() {
	case "single":
		//TODO: equation verification
		if p.method.Hidden {
			p.method.Show()
		}
	}
}

func (p *SingleProblemPanel) openAdditionalInfo(sender *gowd.Element, event *gowd.EventElement) {
	switch sender.GetValue() {
	case "none":
		em["gin"].Hide()
		em["dfin"].Hide()
		em["d2fin"].Hide()
		p.in.Hide()
	case "bisection", "false_pos", "secant":
		em["gin"].Hide()
		em["dfin"].Hide()
		em["d2fin"].Hide()
		p.in.Show()
		em["x1in"].Show()
		em["x2in"].Show()
		em["dxin"].Hide()
		em["tolin"].Show()
		em["itin"].Show()
	case "search":
		em["gin"].Hide()
		em["dfin"].Hide()
		em["d2fin"].Hide()
		p.in.Show()
		em["x1in"].Show()
		em["x2in"].Hide()
		em["dxin"].Show()
		em["tolin"].Hide()
		em["itin"].Show()
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

func (p *SingleProblemPanel) gfuncentry(sender *gowd.Element, event *gowd.EventElement) {
	if sender.GetValue() != "" && em["single_method_selector"].GetValue() == "fixed_point" {
		p.in.Show()
		em["x1in"].Show()
		em["x2in"].Hide()
		em["dxin"].Hide()
		em["tolin"].Show()
		em["itin"].Show()
	}
}

func (p *SingleProblemPanel) dfentry(sender *gowd.Element, event *gowd.EventElement) {
	if val := em["single_method_selector"].GetValue(); sender.GetValue() != "" && (val == "newton" || val == "multi") {
		p.in.Show()
		em["x1in"].Show()
		em["x2in"].Hide()
		em["dxin"].Hide()
		em["tolin"].Show()
		em["itin"].Show()
	}
}

func (p *SingleProblemPanel) d2fentry(sender *gowd.Element, event *gowd.EventElement) {
	if sender.GetValue() != "" && em["single_method_selector"].GetValue() == "multi" && em["d2f"].GetValue() != "" {
		p.in.Show()
		em["x1in"].Show()
		em["x2in"].Hide()
		em["dxin"].Hide()
		em["tolin"].Show()
		em["itin"].Show()
	}
}

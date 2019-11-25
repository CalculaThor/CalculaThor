package gui

import (
	"github.com/dtylman/gowd"
)

type ProblemPanel struct {
	*gowd.Element
	definition *gowd.Element
	problem    *gowd.Element
	method     *gowd.Element
	in         *gowd.Element
}

func (p *ProblemPanel) beginProblem() {
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
	em["submit_button"].Hide()

	em["single_problem_panel"].Hide()
	em["system_problem_panel"].Hide()
	em["solution_panel"].Hide()
	em["interpolation_problem_panel"].Hide()

	em["type_selector"].OnEvent("onchange", p.openProblem)
	em["single_problem_def"].OnEvent("onchange", p.openMethods)
	em["single_method_selector"].OnEvent("onchange", p.openAdditionalInfo)
	em["g"].OnEvent("onchange", p.gFuncEntry)
	em["df"].OnEvent("onchange", p.dfEntry)
	em["d2f"].OnEvent("onchange", p.d2fEntry)

}

func (p *ProblemPanel) openProblem(sender *gowd.Element, event *gowd.EventElement) {
	em["submit_button"].Hide()
	em["solution_panel"].Hide()
	switch sender.GetValue() {
	case "single":
		p.problem.Show()
		em["single_problem_panel"].Show()
		em["system_problem_panel"].Hide()
		em["interpolation_problem_panel"].Hide()
	case "system":
		em["single_problem_panel"].Hide()
		em["system_problem_panel"].Show()
		em["interpolation_problem_panel"].Hide()
		beginSystemProblem()
	case "inter":
		em["single_problem_panel"].Hide()
		em["system_problem_panel"].Hide()
		em["interpolation_problem_panel"].Show()
		beginInterpolationProblem()
	}
}

func (p *ProblemPanel) openMethods(sender *gowd.Element, event *gowd.EventElement) {
	switch em["type_selector"].GetValue() {
	case "single":
		if p.method.Hidden {
			p.method.Show()
		}
	}
}

func (p *ProblemPanel) openAdditionalInfo(sender *gowd.Element, event *gowd.EventElement) {
	em["solution_panel"].Hide()
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
		em["relative"].Show()
		em["absolute"].Show()
		em["submit_button"].Show()
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
		em["relative"].Hide()
		em["absolute"].Hide()
		em["submit_button"].Show()
	case "fixed_point":
		em["gin"].Show()
		em["dfin"].Hide()
		em["d2fin"].Hide()
		p.in.Hide()
		em["submit_button"].Hide()
	case "newton":
		em["gin"].Hide()
		em["dfin"].Show()
		em["d2fin"].Hide()
		p.in.Hide()
		em["submit_button"].Hide()
	case "multi":
		em["gin"].Hide()
		em["dfin"].Show()
		em["d2fin"].Show()
		p.in.Hide()
		em["submit_button"].Hide()

	}
}

func (p *ProblemPanel) gFuncEntry(sender *gowd.Element, event *gowd.EventElement) {
	if sender.GetValue() != "" && em["single_method_selector"].GetValue() == "fixed_point" {
		p.in.Show()
		em["x1in"].Show()
		em["x2in"].Hide()
		em["dxin"].Hide()
		em["tolin"].Show()
		em["itin"].Show()
		em["relative"].Show()
		em["absolute"].Show()
		em["submit_button"].Show()
	}
}

func (p *ProblemPanel) dfEntry(sender *gowd.Element, event *gowd.EventElement) {
	if val := em["single_method_selector"].GetValue(); sender.GetValue() != "" && (val == "newton" || val == "multi") {
		p.in.Show()
		em["x1in"].Show()
		em["x2in"].Hide()
		em["dxin"].Hide()
		em["tolin"].Show()
		em["itin"].Show()
		em["relative"].Show()
		em["absolute"].Show()
		em["submit_button"].Show()
	}
}

func (p *ProblemPanel) d2fEntry(sender *gowd.Element, event *gowd.EventElement) {
	if sender.GetValue() != "" && em["single_method_selector"].GetValue() == "multi" && em["d2f"].GetValue() != "" {
		p.in.Show()
		em["x1in"].Show()
		em["x2in"].Hide()
		em["dxin"].Hide()
		em["tolin"].Show()
		em["itin"].Show()
		em["relative"].Show()
		em["absolute"].Show()
		em["submit_button"].Show()
	}
}

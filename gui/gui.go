package gui

import (
	"github.com/dtylman/gowd"
	"os"
)

var em gowd.ElementsMap

type App struct {
	*gowd.Element
	problem  *singleProblemPanel
	solution *singleSolutionPanel
}

func Hellogui() *App {
	a := &App{}

	fullgui := gowd.NewElement("head")
	file, err := os.Open("gui/Interface.html")
	if err != nil {
		panic(err)
	}
	em = gowd.NewElementMap()
	elements, err := gowd.ParseElements(file, em)
	if err != nil {
		panic(err)
	}
	a.problem = &singleProblemPanel{}
	a.problem.beginProblem()

	a.solution = &singleSolutionPanel{}
	a.solution.beginSolution()

	for _, element := range elements {
		fullgui.AddElement(element)
	}

	a.Element = fullgui

	return a
}

func (a *App) Run() {
	err := gowd.Run(a.Element)
	if err != nil {
		panic(err)
	}
}

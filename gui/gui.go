package gui

import (
	"github.com/dtylman/gowd"
	"os"
)

var em gowd.ElementsMap

type app struct {
	*gowd.Element
	problem *problemPanel
}

func Hellogui() {
	a := &app{}

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
	a.Element = em["app"]
	a.problem = beginPanel()

	for _, element := range elements {
		fullgui.AddElement(element)
	}

	err = gowd.Run(fullgui)
	if err != nil {
		panic(err)
	}
}

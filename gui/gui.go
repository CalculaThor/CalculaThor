package gui

import (
	"github.com/dtylman/gowd"
	"os"
)

func Hellogui() {

	fullgui := gowd.NewElement("head")
	file, err := os.Open("gui/Interface.html")
	if err != nil {
		panic(err)
	}
	elements, err := gowd.ParseElements(file, nil)
	if err != nil {
		panic(err)
	}
	for _, element := range elements {
		fullgui.AddElement(element)
	}

	err = gowd.Run(fullgui)
	if err != nil {
		panic(err)
	}
}

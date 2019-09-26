package gui

import (
	"github.com/dtylman/gowd"
	"os"
)

func Hellogui() {

	fullgui:= gowd.NewElement("body")
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

	gowd.Run(fullgui)

}
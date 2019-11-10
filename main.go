package main

import (
	"calculathor/analyzer"
	"calculathor/gui"
)

func main() {
	app := gui.Hellogui()
	analyzer.BeginAnalyzer(app.Elements)
	app.Run()
}

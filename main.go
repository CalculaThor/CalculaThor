package main

import (
	"github.com/CalculaThor/CalculaThor/analyzer"
	"github.com/CalculaThor/CalculaThor/gui"
)

func main() {
	app := gui.Hellogui()
	analyzer.BeginAnalyzer(app.Elements)
	app.Run()
}

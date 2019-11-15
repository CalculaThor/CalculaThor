package gui

import (
	"fmt"
	"github.com/dtylman/gowd"
	"math/rand"
	"strconv"
)

var numberOfVars int
var aDone, bDone bool

func beginSystemProblem() {
	em["nvarsin"].Show()
	em["matrix_in"].Hide()
	em["se_methods"].Hide()

	em["n_vars"].OnEvent("onchange", setN)

	em["A_button"].OnEvent("onclick", enterMatrix)
	em["B_button"].OnEvent("onclick", enterVector)
}

func setN(sender *gowd.Element, event *gowd.EventElement) {
	n, err := strconv.Atoi(sender.GetValue())
	if err == nil {
		if n < 2 {
			n = 2
		}
		numberOfVars = n
		em["matrix_in"].Show()
	}
}

func enterMatrix(sender *gowd.Element, event *gowd.EventElement) {
	mod := em["matrix_modal"]
	mod.RemoveElements()
	table := gowd.NewElement("table")
	var row, in, elm *gowd.Element
	var id string
	for i := 0; i < numberOfVars; i++ {
		row = gowd.NewElement("tr")
		for j := 0; j < numberOfVars; j++ {
			in = gowd.NewElement("input")
			in.SetAttribute("type", "number")
			id = fmt.Sprintf("A%d_%d", i, j)
			in.SetID(id)
			em[id] = in
			elm = gowd.NewElement("td")
			elm.AddElement(in)
			row.AddElement(elm)
		}
		table.AddElement(row)
	}
	mod.AddElement(table)

	ok := gowd.NewElement("input")
	ok.SetAttribute("type", "button")
	ok.SetID("vector_ok")
	ok.SetValue("Ok")
	ok.OnEvent("onclick", checkMatrix)

	random := gowd.NewElement("input")
	random.SetAttribute("type", "button")
	random.SetID("vector_random")
	random.SetValue("Generate random numbers")
	random.OnEvent("onclick", generateMatrix)

	mod.AddElement(ok)
	mod.AddElement(random)
}

func enterVector(sender *gowd.Element, event *gowd.EventElement) {
	mod := em["vector_modal"]
	mod.RemoveElements()
	table := gowd.NewElement("table")
	var row, in, elm *gowd.Element
	var id string
	for i := 0; i < numberOfVars; i++ {
		row = gowd.NewElement("tr")
		in = gowd.NewElement("input")
		in.SetAttribute("type", "number")
		id = fmt.Sprintf("B%d", i)
		in.SetID(id)
		em[id] = in
		elm = gowd.NewElement("td")
		elm.AddElement(in)
		row.AddElement(elm)
		table.AddElement(row)
	}
	mod.AddElement(table)

	ok := gowd.NewElement("input")
	ok.SetAttribute("type", "button")
	ok.SetID("matrix_ok")
	ok.SetValue("Ok")
	ok.OnEvent("onclick", checkVector)


	random := gowd.NewElement("input")
	random.SetAttribute("type", "button")
	random.SetID("matrix_random")
	random.SetValue("Generate random numbers")
	random.OnEvent("onclick", generateVector)

	mod.AddElement(ok)
	mod.AddElement(random)
}

func checkMatrix(sender *gowd.Element, event *gowd.EventElement) {
	for i := 0; i < numberOfVars; i++ {
		for j := 0; j < numberOfVars; j++ {
			if em[fmt.Sprintf("A%d_%d", i, j)].GetValue() == "" {
				aDone = false
				return
			}
		}
	}
	aDone = true
	if bDone && aDone{
		em["se_methods"].Show()
	}
}

func checkVector(sender *gowd.Element, event *gowd.EventElement) {
	for i := 0; i < numberOfVars; i++ {
		if em[fmt.Sprintf("B%d", i)].GetValue() == "" {
			bDone = false
			return
		}
	}
	bDone = true
	if bDone && aDone{
		em["se_methods"].Show()
	}
}

func generateMatrix(sender *gowd.Element, event *gowd.EventElement) {
	var n float64
	for i := 0; i < numberOfVars; i++ {
		for j := 0; j < numberOfVars; j++ {
			n = rand.Float64() * 200 - 100
			em[fmt.Sprintf("A%d_%d", i, j)].SetValue(fmt.Sprintf("%g", n))
		}
	}
}

func generateVector(sender *gowd.Element, event *gowd.EventElement) {
	var n float64
	for i := 0; i < numberOfVars; i++ {
		n = rand.Float64() * 200 - 100
		em[fmt.Sprintf("B%d", i)].SetValue(fmt.Sprintf("%g", n))
	}
}
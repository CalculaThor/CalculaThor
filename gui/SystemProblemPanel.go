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
	em["system_data"].Hide()

	em["n_vars"].OnEvent("onchange", setN)

	em["A_button"].OnEvent("onclick", enterMatrix)
	em["B_button"].OnEvent("onclick", enterVector)
	em["xi_button"].OnEvent("onclick", initialXEntry)

	em["system_method_selector"].OnEvent("onchange", checkSEMethods)
}

func setN(sender *gowd.Element, event *gowd.EventElement) {
	em["submit_button"].Hide()
	em["solution_panel"].Hide()
	em["se_methods"].Hide()
	em["system_data"].Hide()
	aDone = false
	bDone = false
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
	em["se_methods"].Hide()
	mod := em["matrix_modal"]
	mod.Show()
	mod.RemoveElements()
	table := gowd.NewElement("table")
	table.SetClass("table table-sm table-bordered p-4 table-responsive")
	var row, in, elm *gowd.Element
	var id string
	for i := 0; i < numberOfVars; i++ {
		row = gowd.NewElement("tr")
		for j := 0; j < numberOfVars; j++ {
			in = gowd.NewElement("input")
			in.SetAttribute("type", "number")
			id = fmt.Sprintf("A%d_%d", i, j)
			in.SetID(id)
			in.SetAttribute("placeholder", fmt.Sprintf("A%d-%d", i+1, j+1))
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
	ok.SetClass("btn btn-secondary text-light")
	ok.OnEvent("onclick", checkMatrix)

	random := gowd.NewElement("input")
	random.SetAttribute("type", "button")
	random.SetID("vector_random")
	random.SetValue("Generate random numbers")
	random.SetClass("btn btn-secondary text-light")
	random.OnEvent("onclick", generateMatrix)

	mod.AddElement(ok)
	mod.AddElement(random)
}

func enterVector(sender *gowd.Element, event *gowd.EventElement) {
	em["solution_panel"].Hide()
	em["se_methods"].Hide()
	mod := em["vector_modal"]
	mod.Show()
	mod.RemoveElements()
	table := gowd.NewElement("table")
	table.SetClass("table table-sm table-bordered p-4 table-responsive")
	var row, in, elm *gowd.Element
	var id string
	for i := 0; i < numberOfVars; i++ {
		row = gowd.NewElement("tr")
		in = gowd.NewElement("input")
		in.SetAttribute("type", "number")
		id = fmt.Sprintf("B%d", i)
		in.SetID(id)
		in.SetAttribute("placeholder", fmt.Sprintf("B%d", i+1))
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
	ok.SetClass("btn btn-secondary text-light")
	ok.OnEvent("onclick", checkVector)

	random := gowd.NewElement("input")
	random.SetAttribute("type", "button")
	random.SetID("matrix_random")
	random.SetValue("Generate random numbers")
	random.SetClass("btn btn-secondary text-light")
	random.OnEvent("onclick", generateVector)

	mod.AddElement(ok)
	mod.AddElement(random)
}

func checkMatrix(sender *gowd.Element, event *gowd.EventElement) {
	em["se_methods"].Hide()
	for i := 0; i < numberOfVars; i++ {
		for j := 0; j < numberOfVars; j++ {
			if em[fmt.Sprintf("A%d_%d", i, j)].GetValue() == "" {
				aDone = false
				em["matrix_modal"].Hide()
				return
			}
		}
	}
	aDone = true
	if bDone && aDone {
		em["se_methods"].Show()
	} else {
		em["solution_panel"].Hide()
	}
	em["matrix_modal"].Hide()
}

func checkVector(sender *gowd.Element, event *gowd.EventElement) {
	em["se_methods"].Hide()
	for i := 0; i < numberOfVars; i++ {
		if em[fmt.Sprintf("B%d", i)].GetValue() == "" {
			bDone = false
			em["vector_modal"].Hide()
			return
		}
	}
	bDone = true
	if bDone && aDone {
		em["se_methods"].Show()
	} else {
		em["solution_panel"].Hide()
	}
	em["vector_modal"].Hide()
}

func generateMatrix(sender *gowd.Element, event *gowd.EventElement) {
	var n float64
	for i := 0; i < numberOfVars; i++ {
		for j := 0; j < numberOfVars; j++ {
			n = rand.Float64()*200 - 100
			em[fmt.Sprintf("A%d_%d", i, j)].SetValue(fmt.Sprintf("%g", n))
		}
	}
}

func generateVector(sender *gowd.Element, event *gowd.EventElement) {
	var n float64
	for i := 0; i < numberOfVars; i++ {
		n = rand.Float64()*200 - 100
		em[fmt.Sprintf("B%d", i)].SetValue(fmt.Sprintf("%g", n))
	}
}

func checkSEMethods(sender *gowd.Element, event *gowd.EventElement) {
	em["solution_panel"].Hide()
	em["submit_button"].Show()
	switch sender.GetValue() {
	case "none":
		em["system_data"].Hide()
		em["submit_button"].Hide()
		em["xi_button"].Hide()
	case "gauss", "ppivoting", "tpivoting", "doolittle", "croud", "cholesky":
		em["system_data"].Hide()
		em["SE_tolin"].Hide()
		em["SE_itin"].Hide()
		em["initial_values"].Hide()
		em["w_in"].Hide()
		em["SE_absolute"].Hide()
		em["SE_relative"].Hide()
		em["infinity"].Hide()
		em["euclidean"].Hide()
		em["xi_button"].Hide()
	case "jacobi":
		em["system_data"].Show()
		em["SE_tolin"].Show()
		em["SE_itin"].Show()
		em["w_in"].Hide()
		em["initial_values"].Show()
		em["SE_absolute"].Show()
		em["SE_relative"].Show()
		em["infinity"].Show()
		em["euclidean"].Show()
		em["xi_button"].Show()
		em["submit_button"].Hide()
	case "gauss_seidel":
		em["system_data"].Show()
		em["SE_tolin"].Show()
		em["SE_itin"].Show()
		em["w_in"].Show()
		em["initial_values"].Show()
		em["SE_absolute"].Show()
		em["SE_relative"].Show()
		em["infinity"].Show()
		em["euclidean"].Show()
		em["xi_button"].Show()
		em["submit_button"].Hide()
	default:
		em["system_data"].Hide()
		em["submit_button"].Hide()
	}
}

func initialXEntry(sender *gowd.Element, event *gowd.EventElement) {
	mod := em["initial_values"]
	mod.Show()
	mod.RemoveElements()
	table := gowd.NewElement("table")
	table.SetClass("table table-sm table-bordered p-4 table-responsive")
	var row, in, elm *gowd.Element
	var id string
	for i := 0; i < numberOfVars; i++ {
		row = gowd.NewElement("tr")
		in = gowd.NewElement("input")
		in.SetAttribute("type", "number")
		id = fmt.Sprintf("Xi%d", i)
		in.SetID(id)
		in.SetAttribute("placeholder", fmt.Sprintf("Xi-%d", i+1))
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
	ok.SetClass("btn btn-secondary text-light")
	ok.OnEvent("onclick", checkValues)

	random := gowd.NewElement("input")
	random.SetAttribute("type", "button")
	random.SetID("matrix_random")
	random.SetValue("Generate random numbers")
	random.SetClass("btn btn-secondary text-light")
	random.OnEvent("onclick", generateValues)

	mod.AddElement(ok)
	mod.AddElement(random)
}

func generateValues(sender *gowd.Element, event *gowd.EventElement) {
	var n float64
	for i := 0; i < numberOfVars; i++ {
		n = rand.Float64()*200 - 100
		em[fmt.Sprintf("Xi%d", i)].SetValue(fmt.Sprintf("%g", n))
	}
}

func checkValues(sender *gowd.Element, event *gowd.EventElement) {
	for i := 0; i < numberOfVars; i++ {
		if em[fmt.Sprintf("Xi%d", i)].GetValue() == "" {
			em["initial_values"].Hide()
			em["submit_button"].Hide()
			return
		}
	}
	em["submit_button"].Show()
	em["initial_values"].Hide()
}

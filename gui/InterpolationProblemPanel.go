package gui

import (
	"fmt"
	"github.com/dtylman/gowd"
	"strconv"
)

var nPoints int

func beginInterpolationProblem() {
	em["npointsin"].Show()
	em["points_button"].Hide()
	em["points_in"].Hide()
	em["ip_methods"].Hide()
	em["xip_in"].Hide()

	em["n_points"].OnEvent("onchange", loadNPoints)
	em["points_button"].OnEvent("onclick", loadPoints)
	em["ip_method_selector"].OnEvent("onchange", checkIPMethods)
	em["x_ip"].OnEvent("onchange", checkIPValue)
}

func loadNPoints(sender *gowd.Element, event *gowd.EventElement) {
	em["submit_button"].Hide()
	em["solution_panel"].Hide()
	em["points_in"].Hide()
	em["ip_methods"].Hide()
	em["xip_in"].Hide()
	n, err := strconv.Atoi(sender.GetValue())
	if err != nil {
		return
	}
	nPoints = n
	em["points_button"].Show()
}

func loadPoints(sender *gowd.Element, event *gowd.EventElement) {
	em["ip_methods"].Hide()
	panel := em["points_in"]
	panel.Show()
	panel.RemoveElements()
	table := gowd.NewElement("table")
	table.SetClass("table table-sm table-bordered p-4 table-responsive")

	header := gowd.NewElement("tr")
	xlabel := gowd.NewElement("th")
	xlabel.SetText("X")
	header.AddElement(xlabel)
	ylabel := gowd.NewElement("th")
	ylabel.SetText("Y")
	header.AddElement(ylabel)
	table.AddElement(header)

	var row, in, elm *gowd.Element
	var id string
	for i := 0; i < nPoints; i++ {
		row = gowd.NewElement("tr")

		in = gowd.NewElement("input")
		in.SetAttribute("type", "number")
		id = fmt.Sprintf("X%d", i)
		in.SetID(id)
		in.SetAttribute("placeholder", fmt.Sprintf("X%d", i+1))
		em[id] = in
		elm = gowd.NewElement("td")
		elm.AddElement(in)
		row.AddElement(elm)

		in = gowd.NewElement("input")
		in.SetAttribute("type", "number")
		id = fmt.Sprintf("Y%d", i)
		in.SetID(id)
		in.SetAttribute("placeholder", fmt.Sprintf("Y%d", i+1))
		em[id] = in
		elm = gowd.NewElement("td")
		elm.AddElement(in)
		row.AddElement(elm)

		table.AddElement(row)
	}
	panel.AddElement(table)

	ok := gowd.NewElement("input")
	ok.SetAttribute("type", "button")
	ok.SetID("vector_ok")
	ok.SetValue("Ok")
	ok.SetClass("btn btn-secondary text-light m-3")
	ok.OnEvent("onclick", checkPoints)

	panel.AddElement(ok)
}

func checkPoints(sender *gowd.Element, event *gowd.EventElement) {
	em["ip_methods"].Hide()
	for i := 0; i < nPoints; i++ {
		if em[fmt.Sprintf("X%d", i)].GetValue() == "" || em[fmt.Sprintf("Y%d", i)].GetValue() == "" {
			em["points_in"].Hide()
			return
		}
	}
	em["ip_methods"].Show()
	em["points_in"].Hide()
}

func checkIPMethods(sender *gowd.Element, event *gowd.EventElement) {
	em["solution_panel"].Hide()
	em["submit_button"].Hide()
	if sender.GetValue() == "none" {
		em["xip_in"].Hide()
		return
	}
	em["xip_in"].Show()
	if sender.GetValue() == "newton" {
		em["submit_button"].Show()
	}
}

func checkIPValue(sender *gowd.Element, event *gowd.EventElement) {
	em["solution_panel"].Hide()
	if sender.GetValue() == "" && em["ip_method_selector"].GetValue() != "newton" {
		em["submit_button"].Hide()
		return
	}
	em["submit_button"].Show()
}

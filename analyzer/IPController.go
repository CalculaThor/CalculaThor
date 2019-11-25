package analyzer

import (
	"fmt"
	"github.com/CalculaThor/CalculaThor/analyzer/ipanalyzer"
	"github.com/dtylman/gowd"
	"strconv"
)

var nPoints int
var points []ipanalyzer.Point

func solveInterpolation() {
	em["solution_panel"].RemoveElements()
	n, err := strconv.Atoi(em["n_points"].GetValue())
	if err != nil {
		titl := gowd.NewElement("h4")
		titl.SetText("Bad In")
		em["solution_panel"].AddElement(titl)
		return
	}
	nPoints = n
	err = loadPoints()
	if err != nil {
		titl := gowd.NewElement("h4")
		titl.SetText("Bad In")
		em["solution_panel"].AddElement(titl)
		return
	}
	ipanalyzer.SetPoints(points)

	switch em["ip_method_selector"].GetValue() {
	case "newton":
		ipanalyzer.NewtonDifDiv()
		if ipanalyzer.BadIn() {
			titl := gowd.NewElement("h4")
			titl.SetText("Bad In")
			em["solution_panel"].AddElement(titl)
			return
		}
		table := ipanalyzer.GetNewtonTable()
		showInterpolationTable(table)
		x, err := strconv.ParseFloat(em["x_ip"].GetValue(), 64)
		if err == nil {
			ans := ipanalyzer.InterpolateNewton(x)
			ipans := gowd.NewElement("h4")
			ipans.SetText(fmt.Sprintf("Your f(x) is %g", ans))
			em["solution_panel"].AddElement(ipans)
		}

	case "lagrange":
		ansStr := ipanalyzer.LagrangeString()
		if ipanalyzer.BadIn() {
			titl := gowd.NewElement("h3")
			titl.SetText("Bad In")
			em["solution_panel"].AddElement(titl)
			return
		}
		ansStr = "p(x) = " + ansStr
		x, err := strconv.ParseFloat(em["x_ip"].GetValue(), 64)
		if err != nil {
			titl := gowd.NewElement("h3")
			titl.SetText("Bad In")
			em["solution_panel"].AddElement(titl)
			return
		}
		ans := ipanalyzer.LagrangeInterpolation(x)
		if ipanalyzer.BadIn() {
			titl := gowd.NewElement("h3")
			titl.SetText("Bad In")
			em["solution_panel"].AddElement(titl)
			return
		}
		titl := gowd.NewElement("h3")
		titl.SetText("The polynomial is:")
		em["solution_panel"].AddElement(titl)
		p := gowd.NewElement("p")
		p.SetText(ansStr)
		p.AddElement(gowd.NewElement("br"))
		em["solution_panel"].AddElement(p)
		ipans := gowd.NewElement("h4")
		ipans.SetText(fmt.Sprintf("Your f(x) is %g", ans))
		em["solution_panel"].AddElement(ipans)
	case "neville":
		x, err := strconv.ParseFloat(em["x_ip"].GetValue(), 64)
		if err != nil {
			titl := gowd.NewElement("h3")
			titl.SetText("Bad In")
			em["solution_panel"].AddElement(titl)
			return
		}
		ans := ipanalyzer.NevilleInterpolation(x)
		if ipanalyzer.BadIn() {
			titl := gowd.NewElement("h3")
			titl.SetText("Bad In")
			em["solution_panel"].AddElement(titl)
			return
		}
		table := ipanalyzer.GetNevilleTable()
		showInterpolationTable(table)
		ipans := gowd.NewElement("h4")
		ipans.SetText(fmt.Sprintf("Your f(x) is %g", ans))
		em["solution_panel"].AddElement(ipans)

	default:
		titl := gowd.NewElement("h3")
		titl.SetText("Bad In")
		em["solution_panel"].AddElement(titl)
		return
	}

}

func loadPoints() error {
	points = make([]ipanalyzer.Point, nPoints)
	var x, y float64
	var err error
	for i := 0; i < nPoints; i++ {
		x, err = strconv.ParseFloat(em[fmt.Sprintf("X%d", i)].GetValue(), 64)
		if err != nil {
			return err
		}
		y, err = strconv.ParseFloat(em[fmt.Sprintf("Y%d", i)].GetValue(), 64)
		if err != nil {
			return err
		}
		points[i] = ipanalyzer.Point{x, y}
	}
	return nil
}

func showInterpolationTable(t [][]float64) {
	div := gowd.NewElement("div")
	titl := gowd.NewElement("h3")
	titl.SetText("Table")
	div.AddElement(titl)
	table := gowd.NewElement("table")
	table.SetClass("table table-sm table-bordered p-4 table-responsive")
	header := gowd.NewElement("tr")
	header.SetClass("text-center")
	xn := gowd.NewElement("th")
	xn.SetText("X")
	header.AddElement(xn)
	for i := 0; i < nPoints; i++ {
		tit := gowd.NewElement("th")
		tit.SetText(fmt.Sprintf("#%d", i+1))
		header.AddElement(tit)
	}

	table.AddElement(header)
	for i := 0; i < len(t); i++ {
		row := gowd.NewElement("tr")
		x := gowd.NewElement("td")
		x.SetText(fmt.Sprintf("%g", points[i].X))
		row.AddElement(x)
		for j := 0; j < len(t[i]); j++ {
			el := gowd.NewElement("td")
			el.SetText(fmt.Sprintf("%g", t[i][j]))
			row.AddElement(el)
		}
		table.AddElement(row)
	}

	div.AddElement(table)

	em["solution_panel"].AddElement(div)
}

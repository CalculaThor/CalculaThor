package analyzer

import (
	"fmt"
	"github.com/CalculaThor/CalculaThor/analyzer/svanalyzer"
	"github.com/dtylman/gowd"
	"math"
	"strconv"
)

func solveSingleEquation() {
	f := em["single_problem_def"].GetValue()
	err := svanalyzer.SetF(f)
	if err != nil {
		return
	}

	panel := em["solution_panel"]
	panel.RemoveElements()
	panel.Show()
	switch em["single_method_selector"].GetValue() {
	case "bisection":
		a, err := strconv.ParseFloat(em["x1"].GetValue(), 64)
		if err != nil {
			titl := gowd.NewElement("h3")
			titl.SetText("Bad In")
			em["solution_panel"].AddElement(titl)
			return
		}
		b, err := strconv.ParseFloat(em["x2"].GetValue(), 64)
		if err != nil {
			titl := gowd.NewElement("h3")
			titl.SetText("Bad In")
			em["solution_panel"].AddElement(titl)
			return
		}
		tol, err := strconv.ParseFloat(em["tol"].GetValue(), 64)
		if err != nil {
			titl := gowd.NewElement("h3")
			titl.SetText("Bad In")
			em["solution_panel"].AddElement(titl)
			return
		}
		maxit, err := strconv.ParseInt(em["it"].GetValue(), 10, 32)
		if err != nil {
			titl := gowd.NewElement("h3")
			titl.SetText("Bad In")
			em["solution_panel"].AddElement(titl)
			return
		}
		ans := svanalyzer.Bisection(a, b, math.Abs(tol), uint(math.Abs(float64(maxit))))
		if !ans.BadIn && !ans.IsAlmostRoot && !ans.IsRoot {
			panel.AddElement(gowd.NewElement("h2")).SetText("Not enough iterations.")
			loadBisTable(panel)
		} else if ans.IsRoot {
			panel.AddElement(gowd.NewElement("h2")).SetText("The point x = " + strconv.FormatFloat(ans.Root, 'g', 3, 64) + " is a root.")
			loadBisTable(panel)
		} else if ans.IsAlmostRoot {
			panel.AddElement(gowd.NewElement("h2")).SetText("The point x = " + strconv.FormatFloat(ans.Root, 'g', 3, 64) + " is almost a root.")
			loadBisTable(panel)
		} else if ans.BadIn {
			panel.AddElement(gowd.NewElement("h2")).SetText("Bad In.")
		}

	case "false_pos":
		a, err := strconv.ParseFloat(em["x1"].GetValue(), 64)
		if err != nil {
			titl := gowd.NewElement("h3")
			titl.SetText("Bad In")
			em["solution_panel"].AddElement(titl)
			return
		}
		b, err := strconv.ParseFloat(em["x2"].GetValue(), 64)
		if err != nil {
			titl := gowd.NewElement("h3")
			titl.SetText("Bad In")
			em["solution_panel"].AddElement(titl)
			return
		}
		tol, err := strconv.ParseFloat(em["tol"].GetValue(), 64)
		if err != nil {
			titl := gowd.NewElement("h3")
			titl.SetText("Bad In")
			em["solution_panel"].AddElement(titl)
			return
		}
		maxit, err := strconv.ParseInt(em["it"].GetValue(), 10, 32)
		if err != nil {
			titl := gowd.NewElement("h3")
			titl.SetText("Bad In")
			em["solution_panel"].AddElement(titl)
			return
		}
		ans := svanalyzer.FalsePosition(a, b, math.Abs(tol), uint(math.Abs(float64(maxit))))
		if !ans.BadIn && !ans.IsAlmostRoot && !ans.IsRoot {
			panel.AddElement(gowd.NewElement("h2")).SetText("Not enough iterations.")
			loadFalTable(panel)
		} else if ans.IsRoot {
			panel.AddElement(gowd.NewElement("h2")).SetText("The point x = " + strconv.FormatFloat(ans.Root, 'g', 3, 64) + " is a root.")
			loadFalTable(panel)
		} else if ans.IsAlmostRoot {
			panel.AddElement(gowd.NewElement("h2")).SetText("The point x = " + strconv.FormatFloat(ans.Root, 'g', 3, 64) + " is almost a root.")
			loadFalTable(panel)
		} else if ans.BadIn {
			panel.AddElement(gowd.NewElement("h2")).SetText("Bad In.")
		}
	case "secant":
		a, err := strconv.ParseFloat(em["x1"].GetValue(), 64)
		if err != nil {
			titl := gowd.NewElement("h3")
			titl.SetText("Bad In")
			em["solution_panel"].AddElement(titl)
			return
		}
		b, err := strconv.ParseFloat(em["x2"].GetValue(), 64)
		if err != nil {
			titl := gowd.NewElement("h3")
			titl.SetText("Bad In")
			em["solution_panel"].AddElement(titl)
			return
		}
		tol, err := strconv.ParseFloat(em["tol"].GetValue(), 64)
		if err != nil {
			titl := gowd.NewElement("h3")
			titl.SetText("Bad In")
			em["solution_panel"].AddElement(titl)
			return
		}
		maxit, err := strconv.ParseInt(em["it"].GetValue(), 10, 32)
		if err != nil {
			titl := gowd.NewElement("h3")
			titl.SetText("Bad In")
			em["solution_panel"].AddElement(titl)
			return
		}
		ans := svanalyzer.Secant(a, b, math.Abs(tol), uint(math.Abs(float64(maxit))))
		if !ans.BadIn && !ans.IsAlmostRoot && !ans.IsRoot {
			panel.AddElement(gowd.NewElement("h2")).SetText("Not enough iterations.")
			loadSecTable(panel)
		} else if ans.IsRoot {
			panel.AddElement(gowd.NewElement("h2")).SetText("The point x = " + strconv.FormatFloat(ans.Root, 'g', 3, 64) + " is a root.")
			loadSecTable(panel)
		} else if ans.IsAlmostRoot {
			panel.AddElement(gowd.NewElement("h2")).SetText("The point x = " + strconv.FormatFloat(ans.Root, 'g', 3, 64) + " is almost a root.")
			loadSecTable(panel)
		} else if ans.BadIn {
			panel.AddElement(gowd.NewElement("h2")).SetText("Bad In.")
		}

	case "search":
		x0, err := strconv.ParseFloat(em["x1"].GetValue(), 64)
		if err != nil {
			titl := gowd.NewElement("h3")
			titl.SetText("Bad In")
			em["solution_panel"].AddElement(titl)
			return
		}
		dx, err := strconv.ParseFloat(em["dx"].GetValue(), 64)
		if err != nil {
			titl := gowd.NewElement("h3")
			titl.SetText("Bad In")
			em["solution_panel"].AddElement(titl)
			return
		}
		maxit, err := strconv.ParseInt(em["it"].GetValue(), 10, 32)
		if err != nil {
			titl := gowd.NewElement("h3")
			titl.SetText("Bad In")
			em["solution_panel"].AddElement(titl)
			return
		}
		ans := svanalyzer.IncrementalSearch(x0, dx, uint(math.Abs(float64(maxit))))
		if !ans.IsRange && !ans.IsRoot {
			panel.AddElement(gowd.NewElement("h2")).SetText("Not enough iterations.")
		} else if ans.IsRoot {
			panel.AddElement(gowd.NewElement("h2")).SetText("The point x = " + strconv.FormatFloat(ans.Root, 'g', 3, 64) + " is a root.")
		} else if ans.IsRange {
			panel.AddElement(gowd.NewElement("h2")).SetText("There's a root between " + strconv.FormatFloat(ans.A, 'g', 3, 64) + " and " + strconv.FormatFloat(ans.B, 'E', 3, 64) + ".")
		}
	case "fixed_point":
		x0, err := strconv.ParseFloat(em["x1"].GetValue(), 64)
		if err != nil {
			titl := gowd.NewElement("h3")
			titl.SetText("Bad In")
			em["solution_panel"].AddElement(titl)
			return
		}
		tol, err := strconv.ParseFloat(em["tol"].GetValue(), 64)
		if err != nil {
			titl := gowd.NewElement("h3")
			titl.SetText("Bad In")
			em["solution_panel"].AddElement(titl)
			return
		}
		maxit, err := strconv.ParseInt(em["it"].GetValue(), 10, 32)
		if err != nil {
			titl := gowd.NewElement("h3")
			titl.SetText("Bad In")
			em["solution_panel"].AddElement(titl)
			return
		}
		g := em["g"].GetValue()
		err = svanalyzer.SetG(g)
		if err != nil {
			titl := gowd.NewElement("h3")
			titl.SetText("Bad In")
			em["solution_panel"].AddElement(titl)
			return
		}
		ans := svanalyzer.FixedPoint(x0, math.Abs(tol), uint(math.Abs(float64(maxit))))
		if !ans.BadIn && !ans.IsAlmostRoot && !ans.IsRoot {
			panel.AddElement(gowd.NewElement("h2")).SetText("Not enough iterations.")
			loadFixTable(panel)
		} else if ans.IsRoot {
			panel.AddElement(gowd.NewElement("h2")).SetText("The point x = " + strconv.FormatFloat(ans.Root, 'g', 3, 64) + " is a root.")
			loadFixTable(panel)
		} else if ans.IsAlmostRoot {
			panel.AddElement(gowd.NewElement("h2")).SetText("The point x = " + strconv.FormatFloat(ans.Root, 'g', 3, 64) + " is almost a root.")
			loadFixTable(panel)
		} else if ans.BadIn {
			panel.AddElement(gowd.NewElement("h2")).SetText("Bad In.")
		}
	case "newton":
		x0, err := strconv.ParseFloat(em["x1"].GetValue(), 64)
		if err != nil {
			titl := gowd.NewElement("h3")
			titl.SetText("Bad In")
			em["solution_panel"].AddElement(titl)
			return
		}
		tol, err := strconv.ParseFloat(em["tol"].GetValue(), 64)
		if err != nil {
			titl := gowd.NewElement("h3")
			titl.SetText("Bad In")
			em["solution_panel"].AddElement(titl)
			return
		}
		maxit, err := strconv.ParseInt(em["it"].GetValue(), 10, 32)
		if err != nil {
			titl := gowd.NewElement("h3")
			titl.SetText("Bad In")
			em["solution_panel"].AddElement(titl)
			return
		}
		df := em["df"].GetValue()
		err = svanalyzer.SetDF(df)
		if err != nil {
			return
		}
		ans := svanalyzer.Newton(x0, math.Abs(tol), uint(math.Abs(float64(maxit))))
		if !ans.BadIn && !ans.IsAlmostRoot && !ans.IsRoot {
			panel.AddElement(gowd.NewElement("h2")).SetText("Not enough iterations.")
			loadNewTable(panel)
		} else if ans.IsRoot {
			panel.AddElement(gowd.NewElement("h2")).SetText("The point x = " + strconv.FormatFloat(ans.Root, 'g', 3, 64) + " is a root.")
			loadNewTable(panel)
		} else if ans.IsAlmostRoot {
			panel.AddElement(gowd.NewElement("h2")).SetText("The point x = " + strconv.FormatFloat(ans.Root, 'g', 3, 64) + " is almost a root.")
			loadNewTable(panel)
		} else if ans.BadIn {
			panel.AddElement(gowd.NewElement("h2")).SetText("Bad In.")
		}
	case "multi":
		x0, err := strconv.ParseFloat(em["x1"].GetValue(), 64)
		if err != nil {
			titl := gowd.NewElement("h3")
			titl.SetText("Bad In")
			em["solution_panel"].AddElement(titl)
			return
		}
		tol, err := strconv.ParseFloat(em["tol"].GetValue(), 64)
		if err != nil {
			titl := gowd.NewElement("h3")
			titl.SetText("Bad In")
			em["solution_panel"].AddElement(titl)
			return
		}
		maxit, err := strconv.ParseInt(em["it"].GetValue(), 10, 32)
		if err != nil {
			titl := gowd.NewElement("h3")
			titl.SetText("Bad In")
			em["solution_panel"].AddElement(titl)
			return
		}
		df := em["df"].GetValue()
		err = svanalyzer.SetDF(df)
		if err != nil {
			return
		}
		d2f := em["d2f"].GetValue()
		err = svanalyzer.SetD2F(d2f)
		if err != nil {
			return
		}
		ans := svanalyzer.MultipeRoot(x0, tol, uint(math.Abs(float64(maxit))))
		if !ans.BadIn && !ans.IsAlmostRoot && !ans.IsRoot {
			panel.AddElement(gowd.NewElement("h2")).SetText("Not enough iterations.")
			loadMulTable(panel)
		} else if ans.IsRoot {
			panel.AddElement(gowd.NewElement("h2")).SetText("The point x = " + strconv.FormatFloat(ans.Root, 'g', 3, 64) + " is a root.")
			loadMulTable(panel)
		} else if ans.IsAlmostRoot {
			panel.AddElement(gowd.NewElement("h2")).SetText("The point x = " + strconv.FormatFloat(ans.Root, 'g', 3, 64) + " is almost a root.")
			loadMulTable(panel)
		} else if ans.BadIn {
			panel.AddElement(gowd.NewElement("h2")).SetText("Bad In.")
		}
	}
}

func setAbs(sender *gowd.Element, event *gowd.EventElement) {
	svanalyzer.SetAbsoluteError()
}

func setRel(sender *gowd.Element, event *gowd.EventElement) {
	svanalyzer.SetRelativeError()
}

func loadBisTable(panel *gowd.Element) {
	table := svanalyzer.BisectionTable()
	t := panel.AddElement(gowd.NewElement("table"))
	t.SetClass("table table-sm table-bordered p-4 table-responsive")
	header := t.AddElement(gowd.NewElement("tr"))
	header.AddElement(gowd.NewElement("th")).SetText("n")
	header.AddElement(gowd.NewElement("th")).SetText("xi")
	header.AddElement(gowd.NewElement("th")).SetText("xm")
	header.AddElement(gowd.NewElement("th")).SetText("xs")
	header.AddElement(gowd.NewElement("th")).SetText("f(xm)")
	header.AddElement(gowd.NewElement("th")).SetText("Error")
	for _, reg := range table {
		row := t.AddElement(gowd.NewElement("tr"))
		row.AddElement(gowd.NewElement("td")).SetText(fmt.Sprintf("%d", reg.It))
		row.AddElement(gowd.NewElement("td")).SetText(fmt.Sprintf("%g", reg.Xi))
		row.AddElement(gowd.NewElement("td")).SetText(fmt.Sprintf("%g", reg.Xm))
		row.AddElement(gowd.NewElement("td")).SetText(fmt.Sprintf("%g", reg.Xs))
		row.AddElement(gowd.NewElement("td")).SetText(fmt.Sprintf("%g", reg.Fxm))
		row.AddElement(gowd.NewElement("td")).SetText(fmt.Sprintf("%g", reg.Error))
	}
}

func loadFalTable(panel *gowd.Element) {
	table := svanalyzer.FalsePosTable()
	t := panel.AddElement(gowd.NewElement("table"))
	t.SetClass("table table-sm table-bordered p-4 table-responsive")
	header := t.AddElement(gowd.NewElement("tr"))
	header.AddElement(gowd.NewElement("th")).SetText("n")
	header.AddElement(gowd.NewElement("th")).SetText("xi")
	header.AddElement(gowd.NewElement("th")).SetText("xm")
	header.AddElement(gowd.NewElement("th")).SetText("xs")
	header.AddElement(gowd.NewElement("th")).SetText("f(xm)")
	header.AddElement(gowd.NewElement("th")).SetText("Error")
	for _, reg := range table {
		row := t.AddElement(gowd.NewElement("tr"))
		row.AddElement(gowd.NewElement("td")).SetText(fmt.Sprintf("%d", reg.It))
		row.AddElement(gowd.NewElement("td")).SetText(fmt.Sprintf("%g", reg.Xi))
		row.AddElement(gowd.NewElement("td")).SetText(fmt.Sprintf("%g", reg.Xm))
		row.AddElement(gowd.NewElement("td")).SetText(fmt.Sprintf("%g", reg.Xs))
		row.AddElement(gowd.NewElement("td")).SetText(fmt.Sprintf("%g", reg.Fxm))
		row.AddElement(gowd.NewElement("td")).SetText(fmt.Sprintf("%g", reg.Error))
	}
}

func loadFixTable(panel *gowd.Element) {
	table := svanalyzer.FixedPointTable()
	t := panel.AddElement(gowd.NewElement("table"))
	t.SetClass("table table-sm table-bordered p-4 table-responsive")
	header := t.AddElement(gowd.NewElement("tr"))
	header.AddElement(gowd.NewElement("th")).SetText("n")
	header.AddElement(gowd.NewElement("th")).SetText("xn")
	header.AddElement(gowd.NewElement("th")).SetText("f(xn)")
	header.AddElement(gowd.NewElement("th")).SetText("Error")
	for _, reg := range table {
		row := t.AddElement(gowd.NewElement("tr"))
		row.AddElement(gowd.NewElement("td")).SetText(fmt.Sprintf("%d", reg.It))
		row.AddElement(gowd.NewElement("td")).SetText(fmt.Sprintf("%g", reg.Xm))
		row.AddElement(gowd.NewElement("td")).SetText(fmt.Sprintf("%g", reg.Fxm))
		row.AddElement(gowd.NewElement("td")).SetText(fmt.Sprintf("%g", reg.Error))
	}
}

func loadSecTable(panel *gowd.Element) {
	table := svanalyzer.SecantTable()
	t := panel.AddElement(gowd.NewElement("table"))
	t.SetClass("table table-sm table-bordered p-4 table-responsive")
	header := t.AddElement(gowd.NewElement("tr"))
	header.AddElement(gowd.NewElement("th")).SetText("n")
	header.AddElement(gowd.NewElement("th")).SetText("xn")
	header.AddElement(gowd.NewElement("th")).SetText("f(xn)")
	header.AddElement(gowd.NewElement("th")).SetText("Error")
	for _, reg := range table {
		row := t.AddElement(gowd.NewElement("tr"))
		row.AddElement(gowd.NewElement("td")).SetText(fmt.Sprintf("%d", reg.It))
		row.AddElement(gowd.NewElement("td")).SetText(fmt.Sprintf("%g", reg.Xm))
		row.AddElement(gowd.NewElement("td")).SetText(fmt.Sprintf("%g", reg.Fxm))
		row.AddElement(gowd.NewElement("td")).SetText(fmt.Sprintf("%g", reg.Error))
	}
}

func loadMulTable(panel *gowd.Element) {
	table := svanalyzer.MultipleRootTable()
	t := panel.AddElement(gowd.NewElement("table"))
	t.SetClass("table table-sm table-bordered p-4 table-responsive")
	header := t.AddElement(gowd.NewElement("tr"))
	header.AddElement(gowd.NewElement("th")).SetText("n")
	header.AddElement(gowd.NewElement("th")).SetText("xn")
	header.AddElement(gowd.NewElement("th")).SetText("f(xn)")
	header.AddElement(gowd.NewElement("th")).SetText("f'(xn)")
	header.AddElement(gowd.NewElement("th")).SetText("f\"(xn)")
	header.AddElement(gowd.NewElement("th")).SetText("Error")
	for _, reg := range table {
		row := t.AddElement(gowd.NewElement("tr"))
		row.AddElement(gowd.NewElement("td")).SetText(fmt.Sprintf("%d", reg.It))
		row.AddElement(gowd.NewElement("td")).SetText(fmt.Sprintf("%g", reg.Xm))
		row.AddElement(gowd.NewElement("td")).SetText(fmt.Sprintf("%g", reg.Fxm))
		row.AddElement(gowd.NewElement("td")).SetText(fmt.Sprintf("%g", reg.Dfxm))
		row.AddElement(gowd.NewElement("td")).SetText(fmt.Sprintf("%g", reg.D2fxm))
		row.AddElement(gowd.NewElement("td")).SetText(fmt.Sprintf("%g", reg.Error))
	}
}

func loadNewTable(panel *gowd.Element) {
	table := svanalyzer.NewtonTable()
	t := panel.AddElement(gowd.NewElement("table"))
	t.SetClass("table table-sm table-bordered p-4 table-responsive")
	header := t.AddElement(gowd.NewElement("tr"))
	header.AddElement(gowd.NewElement("th")).SetText("n")
	header.AddElement(gowd.NewElement("th")).SetText("xn")
	header.AddElement(gowd.NewElement("th")).SetText("f(xn)")
	header.AddElement(gowd.NewElement("th")).SetText("f'(xn)")
	header.AddElement(gowd.NewElement("th")).SetText("Error")
	for _, reg := range table {
		row := t.AddElement(gowd.NewElement("tr"))
		row.AddElement(gowd.NewElement("td")).SetText(fmt.Sprintf("%d", reg.It))
		row.AddElement(gowd.NewElement("td")).SetText(fmt.Sprintf("%g", reg.Xm))
		row.AddElement(gowd.NewElement("td")).SetText(fmt.Sprintf("%g", reg.Fxm))
		row.AddElement(gowd.NewElement("td")).SetText(fmt.Sprintf("%g", reg.Dfxm))
		row.AddElement(gowd.NewElement("td")).SetText(fmt.Sprintf("%g", reg.Error))
	}
}

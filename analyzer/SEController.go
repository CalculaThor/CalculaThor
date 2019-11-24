package analyzer

import (
	"fmt"
	"github.com/CalculaThor/CalculaThor/analyzer/seanalyzer"
	"github.com/dtylman/gowd"
	"gonum.org/v1/gonum/mat"
	"strconv"
)

var nVars int
var badIn bool

func solveSystemOfEquations() {
	em["solution_panel"].RemoveElements()
	setN()
	loadMatrix()
	loadVector()

	switch em["system_method_selector"].GetValue() {
	case "gauss":
		ans := seanalyzer.Gauss()
		showResults(ans)
		showStages1(seanalyzer.GetGaussSimpleStages())
	case "ppivoting":
		ans := seanalyzer.GaussPartialPivoting()
		showResults(ans)
		showStages1(seanalyzer.GetGaussPartialStages())
	case "tpivoting":
		ans := seanalyzer.GaussTotalPivoting()
		showResults(ans)
		showStages2(seanalyzer.GetGaussTotalStages())
	case "doolittle":
		ans := seanalyzer.DoolittleFactorization()
		showResults(ans)
		showStages3(seanalyzer.GetDoolittleStages())
	case "croud":
		ans := seanalyzer.CroutFactorization()
		showResults(ans)
		showStages3(seanalyzer.GetCroutStages())
	case "cholesky":
		ans := seanalyzer.CholeskyFactorization()
		showResults(ans)
		showStages3(seanalyzer.GetCholeskyStages())
	case "jacobi":
		tol, err := strconv.ParseFloat(em["SE_tol"].GetValue(), 64)
		if err != nil {
			badIn = true
			return
		}
		it, err := strconv.Atoi(em["SE_it"].GetValue())
		if err != nil {
			badIn = true
			return
		}
		ans, _, _ := seanalyzer.Jacobi(loadInitialValues(), tol, uint(it))
		showResultsSlice(ans)
		showStages4(seanalyzer.GetJacobiTable())
	case "gauss_seidel":
		tol, err := strconv.ParseFloat(em["SE_tol"].GetValue(), 64)
		if err != nil {
			badIn = true
			return
		}
		w, err := strconv.ParseFloat(em["w"].GetValue(), 64)
		if err != nil {
			badIn = true
			return
		}
		it, err := strconv.Atoi(em["SE_it"].GetValue())
		if err != nil {
			badIn = true
			return
		}
		ans, _, _ := seanalyzer.GaussSeidelRelaxed(loadInitialValues(), w, tol, uint(it))
		showResultsSlice(ans)
		showStages4(seanalyzer.GetGaussSeidelTable())
	default:
		em["solution_panel"].Hide()
		em["system_data"].Hide()
		em["submit_button"].Hide()
	}
}

func setN() {
	n, err := strconv.Atoi(em["n_vars"].GetValue())
	if err != nil {
		badIn = true
		return
	}
	nVars = n
}

func setRelNorm(sender *gowd.Element, event *gowd.EventElement) {
	seanalyzer.SetRelative()
}

func setAbsNorm(sender *gowd.Element, event *gowd.EventElement) {
	seanalyzer.SetAbsulute()
}

func setInfinity(sender *gowd.Element, event *gowd.EventElement) {
	seanalyzer.SetInfinity()
}

func setEuclidean(sender *gowd.Element, event *gowd.EventElement) {
	seanalyzer.SetEuclidean()
}

func loadMatrix() {
	ans := make([]float64, nVars*nVars)
	var index int
	for i := 0; i < nVars; i++ {
		for j := 0; j < nVars; j++ {
			index = j + i*nVars
			val, err := strconv.ParseFloat(em[fmt.Sprintf("A%d_%d", i, j)].GetValue(), 64)
			if err != nil {
				badIn = true
				return
			}
			ans[index] = val
		}
	}

	seanalyzer.SetCoefficientMatrix(nVars, ans)
}

func loadVector() {
	ans := make([]float64, nVars)
	for i := 0; i < nVars; i++ {
		val, err := strconv.ParseFloat(em[fmt.Sprintf("B%d", i)].GetValue(), 64)
		if err != nil {
			badIn = true
			return
		}
		ans[i] = val
	}

	seanalyzer.SetIndependentTerms(ans)
}

func loadInitialValues() []float64 {
	ans := make([]float64, nVars)
	for i := 0; i < nVars; i++ {
		val, err := strconv.ParseFloat(em[fmt.Sprintf("Xi%d", i)].GetValue(), 64)
		if err != nil {
			badIn = true
			return nil
		}
		ans[i] = val
	}

	return ans
}

func showResults(ans *mat.VecDense) {
	titl := gowd.NewElement("h2")
	titl.SetText("Solution")
	em["solution_panel"].AddElement(titl)
	table := gowd.NewElement("table")
	table.SetClass("table table-sm table-bordered table-responsive")
	header := gowd.NewElement("tr")
	for i := 0; i < nVars; i++ {
		el := gowd.NewElement("th")
		el.SetText(fmt.Sprintf("X %d", i+1))
		header.AddElement(el)
	}

	table.AddElement(header)
	values := gowd.NewElement("tr")
	for i := 0; i < nVars; i++ {
		el := gowd.NewElement("td")
		el.SetText(fmt.Sprintf("%g", ans.AtVec(i)))
		values.AddElement(el)
	}

	table.AddElement(values)
	em["solution_panel"].AddElement(table)
}

func showResultsSlice(ans []float64) {
	table := gowd.NewElement("table")
	table.SetClass("table table-sm table-bordered table-responsive")
	header := gowd.NewElement("tr")
	for i := 0; i < nVars; i++ {
		el := gowd.NewElement("th")
		el.SetText(fmt.Sprintf("X %d", i+1))
		header.AddElement(el)
	}

	table.AddElement(header)
	values := gowd.NewElement("tr")
	for i := 0; i < nVars; i++ {
		el := gowd.NewElement("td")
		el.SetText(fmt.Sprintf("%g", ans[i]))
		values.AddElement(el)
	}

	table.AddElement(values)
	em["solution_panel"].AddElement(table)
}

func showStages1(stages []seanalyzer.Reg1) {
	stagesDiv := gowd.NewElement("div")
	titl := gowd.NewElement("h2")
	titl.SetText("Stages")
	stagesDiv.AddElement(titl)
	for i, stage := range stages {
		div := gowd.NewElement("div")
		div.SetClass("container p-4")
		tit := gowd.NewElement("h3")
		tit.SetText("Stage " + strconv.Itoa(i))
		div.AddElement(tit)

		tit1 := gowd.NewElement("h4")
		tit1.SetText("Matrix")
		div.AddElement(tit1)

		matr := gowd.NewElement("table")
		matr.SetClass("table table-sm table-bordered p-4 table-responsive")
		r, c := stage.Mat.Caps()
		for x := 0; x < r; x++ {
			row := gowd.NewElement("tr")
			for y := 0; y < c; y++ {
				el := gowd.NewElement("td")
				el.SetText(fmt.Sprintf("%g", stage.Mat.At(x, y)))
				row.AddElement(el)
			}
			matr.AddElement(row)
		}
		div.AddElement(matr)

		tit2 := gowd.NewElement("h4")
		tit2.SetText("Multiplicators")
		div.AddElement(tit2)

		mults := gowd.NewElement("table")
		mults.SetClass("table table-sm table-bordered p-4 table-responsive")
		header := gowd.NewElement("tr")
		for index, _ := range stage.Mul {
			el := gowd.NewElement("th")
			el.SetText(fmt.Sprintf("Mul %d, %d", index.I, index.J))
			header.AddElement(el)
		}
		mults.AddElement(header)
		multRow := gowd.NewElement("tr")
		for _, m := range stage.Mul {
			el := gowd.NewElement("td")
			el.SetText(fmt.Sprintf("%g", m))
			multRow.AddElement(el)
		}
		mults.AddElement(multRow)

		div.AddElement(mults)

		stagesDiv.AddElement(div)
	}
	em["solution_panel"].AddElement(stagesDiv)

}

func showStages2(stages []seanalyzer.Reg2) {
	stagesDiv := gowd.NewElement("div")
	titl := gowd.NewElement("h2")
	titl.SetText("Stages")
	stagesDiv.AddElement(titl)
	for i, stage := range stages {
		div := gowd.NewElement("div")
		div.SetClass("container p-4")
		tit := gowd.NewElement("h3")
		tit.SetText("Stage " + strconv.Itoa(i))
		div.AddElement(tit)

		tit1 := gowd.NewElement("h4")
		tit1.SetText("Matrix")
		div.AddElement(tit1)

		matr := gowd.NewElement("table")
		matr.SetClass("table table-sm table-bordered p-4 table-responsive")
		r, c := stage.Mat.Caps()
		for x := 0; x < r; x++ {
			row := gowd.NewElement("tr")
			for y := 0; y < c; y++ {
				el := gowd.NewElement("td")
				el.SetText(fmt.Sprintf("%g", stage.Mat.At(x, y)))
				row.AddElement(el)
			}
			matr.AddElement(row)
		}
		div.AddElement(matr)

		tit2 := gowd.NewElement("h4")
		tit2.SetText("Multiplicators")
		div.AddElement(tit2)

		mults := gowd.NewElement("table")
		mults.SetClass("table table-sm table-bordered p-4 table-responsive")
		header := gowd.NewElement("tr")
		for index, _ := range stage.Mul {
			el := gowd.NewElement("th")
			el.SetText(fmt.Sprintf("Mul %d, %d", index.I, index.J))
			header.AddElement(el)
		}
		mults.AddElement(header)
		multRow := gowd.NewElement("tr")
		for _, m := range stage.Mul {
			el := gowd.NewElement("td")
			el.SetText(fmt.Sprintf("%g", m))
			multRow.AddElement(el)
		}
		mults.AddElement(multRow)

		div.AddElement(mults)

		tit3 := gowd.NewElement("h4")
		tit3.SetText("Marks")
		div.AddElement(tit3)

		ind := gowd.NewElement("table")
		ind.SetClass("table table-sm table-bordered p-4 table-responsive")
		indRow := gowd.NewElement("tr")
		for _, n := range stage.Marks {
			el := gowd.NewElement("td")
			el.SetText(fmt.Sprintf("%d", n))
			indRow.AddElement(el)
		}
		div.AddElement(ind)

		stagesDiv.AddElement(div)
	}
	em["solution_panel"].AddElement(stagesDiv)

}

func showStages3(stages []seanalyzer.Reg3) {
	stagesDiv := gowd.NewElement("div")
	titl := gowd.NewElement("h2")
	titl.SetText("Stages")
	stagesDiv.AddElement(titl)
	for i, stage := range stages {
		div := gowd.NewElement("div")
		div.SetClass("container p-4")
		tit := gowd.NewElement("h3")
		tit.SetText("Stage " + strconv.Itoa(i))
		div.AddElement(tit)

		tit1 := gowd.NewElement("h4")
		tit1.SetText("L Matrix")
		div.AddElement(tit1)

		matr := gowd.NewElement("table")
		matr.SetClass("table table-sm table-bordered p-4 table-responsive")
		r, c := stage.L.Caps()
		for x := 0; x < r; x++ {
			row := gowd.NewElement("tr")
			for y := 0; y < c; y++ {
				el := gowd.NewElement("td")
				el.SetText(fmt.Sprintf("%g", stage.L.At(x, y)))
				row.AddElement(el)
			}
			matr.AddElement(row)
		}
		div.AddElement(matr)

		tit2 := gowd.NewElement("h4")
		tit2.SetText("U Matrix")
		div.AddElement(tit2)

		matr2 := gowd.NewElement("table")
		matr2.SetClass("table table-sm table-bordered p-4 table-responsive")
		r2, c2 := stage.L.Caps()
		for x := 0; x < r2; x++ {
			row := gowd.NewElement("tr")
			for y := 0; y < c2; y++ {
				el := gowd.NewElement("td")
				el.SetText(fmt.Sprintf("%g", stage.U.At(x, y)))
				row.AddElement(el)
			}
			matr2.AddElement(row)
		}
		div.AddElement(matr2)

		stagesDiv.AddElement(div)
	}
	em["solution_panel"].AddElement(stagesDiv)

}

func showStages4(stages []seanalyzer.Reg4) {
	stagesDiv := gowd.NewElement("div")
	titl := gowd.NewElement("h2")
	titl.SetText("Table")
	stagesDiv.AddElement(titl)

	table := gowd.NewElement("table")
	table.SetClass("table table-sm table-bordered p-4 table-responsive")
	header := gowd.NewElement("tr")
	itEl := gowd.NewElement("th")
	itEl.SetText("Iteration")
	header.AddElement(itEl)

	for i := 1; i <= nVars; i++ {
		el := gowd.NewElement("th")
		el.SetText(fmt.Sprintf("X %d", i))
		header.AddElement(el)
	}

	dispEl := gowd.NewElement("th")
	dispEl.SetText("Dispersion")
	header.AddElement(dispEl)
	table.AddElement(header)

	for _, row := range stages {
		rowEl := gowd.NewElement("tr")
		nEl := gowd.NewElement("td")
		nEl.SetText(fmt.Sprintf("%d", row.It))
		rowEl.AddElement(nEl)

		for i := 0; i < len(row.X); i++ {
			el := gowd.NewElement("td")
			el.SetText(fmt.Sprintf("%g", row.X[i]))
			rowEl.AddElement(el)
		}

		disp := gowd.NewElement("td")
		disp.SetText(fmt.Sprintf("%g", row.Disp))
		rowEl.AddElement(disp)
		table.AddElement(rowEl)
	}

	stagesDiv.AddElement(table)
	em["solution_panel"].AddElement(stagesDiv)

}

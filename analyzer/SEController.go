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
	case "tpivoting":
		el := gowd.NewElement("div")
		ans := seanalyzer.GaussTotalPivoting()
		for i := 0; i < ans.Cap(); i++ {
			el1 := gowd.NewElement("h2")
			el1.SetText(strconv.FormatFloat(ans.AtVec(i), 'f', 5, 64))
			el.AddElement(el1)
		}
		em["solution_panel"].AddElement(el)
	case "doolittle":
		ans := seanalyzer.DoolittleFactorization()
		showResults(ans)
	case "croud":
		ans := seanalyzer.CroutFactorization()
		showResults(ans)
	case "cholesky":
		ans := seanalyzer.CholeskyFactorization()
		showResults(ans)
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
	table := gowd.NewElement("table")
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
	titl := gowd.NewElement("h3")
	titl.SetText("Stages")
	for i, stage := range stages {
		div := gowd.NewElement("div")
		tit := gowd.NewElement("h3")
		tit.SetText("Stage " + strconv.Itoa(i))
		div.AddElement(tit)
		matr := gowd.NewElement("table")
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

		stagesDiv.AddElement(div)
	}
	em["solution_panel"].AddElement(stagesDiv)

}

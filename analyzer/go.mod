module github.com/CalculaThor/CalculaThor/analyzer

require (
	github.com/CalculaThor/CalculaThor/analyzer/seanalyzer v0.0.0-20191124094452-0444e7f960e2
	github.com/CalculaThor/CalculaThor/analyzer/svanalyzer v0.0.0-20191111040314-28a9366a5447
	github.com/dtylman/gowd v0.0.0-20190619113956-15e38debca22
)

replace github.com/CalculaThor/CalculaThor/analyzer/svanalyzer => ./svanalyzer

replace github.com/CalculaThor/CalculaThor/analyzer/seanalyzer => ./seanalyzer

go 1.13

module github.com/CalculaThor/CalculaThor/analyzer

require (
	github.com/CalculaThor/CalculaThor/analyzer/ipanalyzer v0.0.0-20191125054046-9868a48780cd // indirect
	github.com/CalculaThor/CalculaThor/analyzer/seanalyzer v0.0.0-20191124094452-0444e7f960e2
	github.com/CalculaThor/CalculaThor/analyzer/svanalyzer v0.0.0-20200415203310-aa37ae6e9bf1
	github.com/dtylman/gowd v0.0.0-20190619113956-15e38debca22
)

replace github.com/CalculaThor/CalculaThor/analyzer/svanalyzer => ./svanalyzer

replace github.com/CalculaThor/CalculaThor/analyzer/seanalyzer => ./seanalyzer

replace github.com/CalculaThor/CalculaThor/analyzer/ipanalyzer => ./ipanalyzer

go 1.13

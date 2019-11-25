module github.com/CalculaThor/CalculaThor

go 1.13

require (
	github.com/CalculaThor/CalculaThor/analyzer v0.0.0-20191110232331-384c76640fc2
	github.com/CalculaThor/CalculaThor/analyzer/ipanalyzer v0.0.0-20191125054046-9868a48780cd // indirect
	github.com/CalculaThor/CalculaThor/gui v0.0.0-20191110215117-1b851d185092
)

replace github.com/CalculaThor/CalculaThor/analyzer => ./analyzer

replace github.com/CalculaThor/CalculaThor/gui => ./gui

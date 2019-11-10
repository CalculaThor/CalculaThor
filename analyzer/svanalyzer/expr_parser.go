package svanalyzer

import (
	"github.com/Knetic/govaluate"
	"math"
)

var f, df, d2f, g *govaluate.EvaluableExpression

func F(x float64) (float64, error) {
	params := make(map[string]interface{}, 3)
	params["x"] = x
	params["e"] = math.E
	params["pi"] = math.Pi
	ans, err := f.Evaluate(params)
	return ans.(float64), err
}

func DF(x float64) (float64, error) {
	params := make(map[string]interface{}, 3)
	params["x"] = x
	params["e"] = math.E
	params["pi"] = math.Pi
	ans, err := df.Evaluate(params)
	return ans.(float64), err
}

func D2F(x float64) (float64, error) {
	params := make(map[string]interface{}, 3)
	params["x"] = x
	params["e"] = math.E
	params["pi"] = math.Pi
	ans, err := d2f.Evaluate(params)
	return ans.(float64), err
}

func G(x float64) (float64, error) {
	params := make(map[string]interface{}, 3)
	params["x"] = x
	params["e"] = math.E
	params["pi"] = math.Pi
	ans, err := g.Evaluate(params)
	return ans.(float64), err
}

func SetF(s string) error {
	functions := map[string]govaluate.ExpressionFunction{
		"sqrt": func(args ...interface{}) (interface{}, error) {
			return math.Sqrt(args[0].(float64)), nil
		},
		"sin": func(args ...interface{}) (interface{}, error) {
			return math.Sin(args[0].(float64)), nil
		},
		"cos": func(args ...interface{}) (interface{}, error) {
			return math.Cos(args[0].(float64)), nil
		},
		"ln": func(args ...interface{}) (interface{}, error) {
			return math.Log(args[0].(float64)), nil
		},
		"log2": func(args ...interface{}) (interface{}, error) {
			return math.Log2(args[0].(float64)), nil
		},
		"log10": func(args ...interface{}) (interface{}, error) {
			return math.Log10(args[0].(float64)), nil
		},
		"exp": func(args ...interface{}) (interface{}, error) {
			return math.Exp(args[0].(float64)), nil
		},
	}
	expr, err := govaluate.NewEvaluableExpressionWithFunctions(s, functions)
	f = expr
	return err
}

func SetDF(s string) error {
	functions := map[string]govaluate.ExpressionFunction{
		"sqrt": func(args ...interface{}) (interface{}, error) {
			return math.Sqrt(args[0].(float64)), nil
		},
		"sin": func(args ...interface{}) (interface{}, error) {
			return math.Sin(args[0].(float64)), nil
		},
		"cos": func(args ...interface{}) (interface{}, error) {
			return math.Cos(args[0].(float64)), nil
		},
		"ln": func(args ...interface{}) (interface{}, error) {
			return math.Log(args[0].(float64)), nil
		},
		"log2": func(args ...interface{}) (interface{}, error) {
			return math.Log2(args[0].(float64)), nil
		},
		"log10": func(args ...interface{}) (interface{}, error) {
			return math.Log10(args[0].(float64)), nil
		},
		"exp": func(args ...interface{}) (interface{}, error) {
			return math.Exp(args[0].(float64)), nil
		},
	}
	expr, err := govaluate.NewEvaluableExpressionWithFunctions(s, functions)
	df = expr
	return err
}

func SetD2F(s string) error {
	functions := map[string]govaluate.ExpressionFunction{
		"sqrt": func(args ...interface{}) (interface{}, error) {
			return math.Sqrt(args[0].(float64)), nil
		},
		"sin": func(args ...interface{}) (interface{}, error) {
			return math.Sin(args[0].(float64)), nil
		},
		"cos": func(args ...interface{}) (interface{}, error) {
			return math.Cos(args[0].(float64)), nil
		},
		"ln": func(args ...interface{}) (interface{}, error) {
			return math.Log(args[0].(float64)), nil
		},
		"log2": func(args ...interface{}) (interface{}, error) {
			return math.Log2(args[0].(float64)), nil
		},
		"log10": func(args ...interface{}) (interface{}, error) {
			return math.Log10(args[0].(float64)), nil
		},
		"exp": func(args ...interface{}) (interface{}, error) {
			return math.Exp(args[0].(float64)), nil
		},
	}
	expr, err := govaluate.NewEvaluableExpressionWithFunctions(s, functions)
	d2f = expr
	return err
}

func SetG(s string) error {
	functions := map[string]govaluate.ExpressionFunction{
		"sqrt": func(args ...interface{}) (interface{}, error) {
			return math.Sqrt(args[0].(float64)), nil
		},
		"sin": func(args ...interface{}) (interface{}, error) {
			return math.Sin(args[0].(float64)), nil
		},
		"cos": func(args ...interface{}) (interface{}, error) {
			return math.Cos(args[0].(float64)), nil
		},
		"ln": func(args ...interface{}) (interface{}, error) {
			return math.Log(args[0].(float64)), nil
		},
		"log2": func(args ...interface{}) (interface{}, error) {
			return math.Log2(args[0].(float64)), nil
		},
		"log10": func(args ...interface{}) (interface{}, error) {
			return math.Log10(args[0].(float64)), nil
		},
		"exp": func(args ...interface{}) (interface{}, error) {
			return math.Exp(args[0].(float64)), nil
		},
	}
	expr, err := govaluate.NewEvaluableExpressionWithFunctions(s, functions)
	g = expr
	return err
}

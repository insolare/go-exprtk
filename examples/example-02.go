package main

import (
	"fmt"
	"math"

	"github.com/insolare/go-exprtk"
)

func example02() {
	var eqn string
	eqn = "if (eqn == 'avg') avg(10, 20); "
	eqn += "else if (eqn == 'max') println('max'); "
	eqn += "else if (eqn == 'min') min(x); "
	eqn += "else if (eqn == 'sum') sum(x); "
	eqn += "else 0; "

	var eqnStr string
	var array []float64 = []float64{1, 2, 3, -4.3, 10, -6.5, 7, 8, -1.3}

	exprtkObj := exprtk.NewExprtk()
	defer exprtkObj.Delete()

	exprtkObj.SetExpression(eqn)
	exprtkObj.AddStringVariable("eqn")
	exprtkObj.AddVectorVariable("x")
	err := exprtkObj.CompileExpression()
	fmt.Println("Err", err)
	exprtkObj.SetVectorVariableValue("x", array)

	eqnStr = "avg"
	exprtkObj.SetStringVariableValue("eqn", eqnStr)
	fmt.Println(math.Round(exprtkObj.GetEvaluatedValue()*100) / 100)

	eqnStr = "max"
	exprtkObj.SetStringVariableValue("eqn", eqnStr)
	fmt.Println(exprtkObj.GetEvaluatedValue())

	eqnStr = "min"
	exprtkObj.SetStringVariableValue("eqn", eqnStr)
	fmt.Println(exprtkObj.GetEvaluatedValue())

	eqnStr = "sum"
	exprtkObj.SetStringVariableValue("eqn", eqnStr)
	fmt.Println(exprtkObj.GetEvaluatedValue())

	eqnStr = "xyz"
	exprtkObj.SetStringVariableValue("eqn", eqnStr)
	fmt.Println(exprtkObj.GetEvaluatedValue())
}

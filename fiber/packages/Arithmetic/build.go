package arithmetic

import (
	"fmt"

	"github.com/lxn/walk/declarative"
)

// Build Return the right databinder & the right template for a flow instruction
func Build(function string) (interface{}, []declarative.Widget) {
	switch function {
	case "Divide":
		return new(ArithmeticDatabinder), DivideTemplate
	case "Multiply":
		return new(ArithmeticDatabinder), MultiplyTemplate
	case "Pow":
		return new(ArithmeticDatabinder), PowTemplate
	case "Sqrt":
		return new(SqrtDatabinder), SqrtTemplate
	case "Substract":
		return new(ArithmeticDatabinder), SubstractTemplate
	case "Sum":
		return new(ArithmeticDatabinder), SumTemplate
	default:
		fmt.Println("GOTOMATE ERROR: Unable to find the function for instruction building")
		return nil, nil
	}
}

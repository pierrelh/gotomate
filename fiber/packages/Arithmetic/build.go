package arithmetic

import "gotomate-astilectron/fiber/template"

// Build Return the right databinder & the right template for a flow instruction
func Build(function string) (interface{}, *template.InstructionTemplate) {
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
		return nil, nil
	}
}

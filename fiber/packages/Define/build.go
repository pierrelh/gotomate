package define

import "gotomate/fiber/template"

// Build Return the right databinder & the right template for a flow instruction
func Build(function string) (interface{}, *template.InstructionTemplate) {
	switch function {
	case "ArrayOfBool":
		return new(ArrayDatabinder), ArrayTemplate
	case "ArrayOfFloat":
		return new(ArrayDatabinder), ArrayTemplate
	case "ArrayOfInt":
		return new(ArrayDatabinder), ArrayTemplate
	case "ArrayOfString":
		return new(ArrayDatabinder), ArrayTemplate
	case "Bool":
		return new(BoolDatabinder), BoolTemplate
	case "Float":
		return new(FloatDatabinder), FloatTemplate
	case "Int":
		return new(IntDatabinder), IntTemplate
	case "String":
		return new(StringDatabinder), StringTemplate
	}
	return nil, nil
}

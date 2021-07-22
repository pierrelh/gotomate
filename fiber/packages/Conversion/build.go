package conversion

import "gotomate/fiber/template"

// Build Return the right databinder & the right template for a clipboard instruction
func Build(function string) (interface{}, *template.InstructionTemplate) {
	switch function {
	case "BoolToFloat":
		return new(BoolConversionDatabinder), BoolConversionTemplate
	case "BoolToInt":
		return new(BoolConversionDatabinder), BoolConversionTemplate
	case "BoolToString":
		return new(BoolConversionDatabinder), BoolConversionTemplate
	case "FloatToInt":
		return new(FloatConversionDatabinder), FloatConversionTemplate
	case "FloatToString":
		return new(FloatConversionDatabinder), FloatConversionTemplate
	case "IntToFloat":
		return new(IntConversionDatabinder), IntConversionTemplate
	case "IntToString":
		return new(IntConversionDatabinder), IntConversionTemplate
	case "StringToBool":
		return new(StringConversionDatabinder), StringConversionTemplate
	case "StringToFloat":
		return new(StringConversionDatabinder), StringConversionTemplate
	case "StringToInt":
		return new(StringConversionDatabinder), StringConversionTemplate
	}
	return nil, nil
}

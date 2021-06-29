package conversion

import (
	"fmt"

	"github.com/lxn/walk/declarative"
)

// Build Return the right databinder & the right template for a clipboard instruction
func Build(function string) (interface{}, []declarative.Widget) {
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
	fmt.Println("GOTOMATE ERROR: Unable to find the function for instruction building")
	return nil, nil
}

package define

import (
	"fmt"

	"github.com/lxn/walk/declarative"
)

// Build Return the right databinder & the right template for a flow instruction
func Build(function string) (interface{}, []declarative.Widget) {
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
	fmt.Println("GOTOMATE ERROR: Unable to find the function for instruction building")
	return nil, nil
}

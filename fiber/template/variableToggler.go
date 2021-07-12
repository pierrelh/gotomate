package template

import "reflect"

// VariableToggler Define a new variableToggler in a Template's Field
type VariableToggler struct {
	Bind     string // The name of the databinder entree to bind the raw value of the VariableToggler
	Checked  bool   // Set the checked status of the VariableToggler, default false
	Disabled bool   // Set the disabled status of the VariableToggler, default false
}

func newVariableToggler(f interface{}) map[string]interface{} {
	var varToggler = map[string]interface{}{}
	varTogReflect := reflect.ValueOf(f)
	if val := varTogReflect.FieldByName("Bind").Interface().(string); val != "" {
		varToggler["Bind"] = val
	}
	if val := varTogReflect.FieldByName("Disabled").Interface().(bool); val {
		varToggler["Disabled"] = val
	}

	varToggler["Checked"] = varTogReflect.FieldByName("Checked").Interface().(bool)
	varToggler["ElementType"] = TypeCheckbox
	return varToggler
}

package template

import "reflect"

// Label Define a new label in a Template's Field
type Label struct {
	Text      string // The text contained in the Label
	Optionnal bool   // Set if the following input is optionnal
}

func newLabel(f interface{}) map[string]interface{} {
	var label = map[string]interface{}{}
	lreflect := reflect.ValueOf(f)
	if val := lreflect.FieldByName("Text").Interface().(string); val != "" {
		label["Text"] = val
	}
	if val := lreflect.FieldByName("Optionnal").Interface().(bool); val {
		label["Optionnal"] = val
	}
	label["ElementType"] = TypeLabel
	return label
}

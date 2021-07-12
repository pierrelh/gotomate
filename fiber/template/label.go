package template

import "reflect"

// Label Define a new label in a Template's Field
type Label struct {
	Text string // The text contained in the Label
}

func newLabel(f interface{}) map[string]interface{} {
	var label = map[string]interface{}{}
	lreflect := reflect.ValueOf(f)
	if val := lreflect.FieldByName("Text").Interface().(string); val != "" {
		label["Text"] = val
	}
	label["ElementType"] = TypeLabel
	return label
}

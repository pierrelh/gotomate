package template

type ElementType string

const (
	TypeCheckbox ElementType = "checkbox"
	TypeInput    ElementType = "input"
	TypeLabel    ElementType = "label"
	TypeOption   ElementType = "option"
	TypeSelect   ElementType = "select"
	TypeTextArea ElementType = "textarea"
	TypeTextView ElementType = "p"
)

type Template []map[string]map[string]interface{}

// Field Define a field of a Template
type Field struct {
	Label           interface{} // The Label of the Field, use template.Label{}.NewElement()
	TextView        interface{} // The TextView of the Field, use template.TextView{}.NewElement()
	Input           interface{} // The Input of the Field, use template.Input{}.NewElement(), template.TextArea{}.NewElement(), or template.Select{}.NewElement()
	VariableToggler interface{} // The VariableToggler of the Field, use template.VariableToggler{}.NewElement()
}

// Build Create a new Template's Field
func (f Field) Build() map[string]map[string]interface{} {
	var field = map[string]map[string]interface{}{}
	if f.Label != nil {
		field["Label"] = newLabel(f.Label)
	}
	if f.TextView != nil {
		field["TextView"] = newTextView(f.TextView)
	}
	if f.Input != nil {
		field["Input"] = createInput(f.Input)
	}
	if f.VariableToggler != nil {
		field["VariableToggler"] = newVariableToggler(f.VariableToggler)
	}
	return field
}

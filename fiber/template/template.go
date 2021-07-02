package template

import (
	"reflect"
)

type Type string
type FontWeight string
type elementType string

const (
	CheckBox      Type = "checkbox"
	Color         Type = "color"
	Date          Type = "date"
	DatetimeLocal Type = "datetime-local"
	Email         Type = "email"
	File          Type = "file"
	Image         Type = "image"
	Month         Type = "month"
	Number        Type = "number"
	Password      Type = "password"
	Tel           Type = "tel"
	Text          Type = "text"
	Url           Type = "url"
	Week          Type = "week"

	TypeLabel    elementType = "label"
	TypeTextView elementType = "p"
	TypeInput    elementType = "input"
	TypeTextArea elementType = "textarea"
	TypeSelect   elementType = "select"
	TypeOption   elementType = "option"
	TypeCheckbox elementType = "checkbox"

	FWNormal FontWeight = "normal"
	FWBold   FontWeight = "bold"
	FW100    FontWeight = "100"
	FW200    FontWeight = "200"
	FW300    FontWeight = "300"
	FW400    FontWeight = "400"
	FW500    FontWeight = "500"
	FW600    FontWeight = "600"
	FW700    FontWeight = "700"
	FW800    FontWeight = "800"
	FW900    FontWeight = "900"
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
		var label = map[string]interface{}{}
		lreflect := reflect.ValueOf(f.Label)
		label["AssignTo"] = lreflect.FieldByName("AssignTo").Interface().(string)
		label["Text"] = lreflect.FieldByName("Text").Interface().(string)
		label["ElementType"] = lreflect.FieldByName("ElementType").Interface().(elementType)
		field["Label"] = label
	}
	if f.TextView != nil {
		var textView = map[string]interface{}{}
		tvreflect := reflect.ValueOf(f.TextView)
		textView["FontWeight"] = tvreflect.FieldByName("FontWeight").Interface().(string)
		textView["Text"] = tvreflect.FieldByName("Text").Interface().(string)
		textView["ElementType"] = tvreflect.FieldByName("ElementType").Interface().(elementType)
		field["TextView"] = textView
	}
	if f.Input != nil {
		var input = map[string]interface{}{}
		inputReflect := reflect.ValueOf(f.Input)
		input["Bind"] = inputReflect.FieldByName("Bind").Interface().(string)
		input["BindVariable"] = inputReflect.FieldByName("BindVariable").Interface().(string)
		input["ID"] = inputReflect.FieldByName("ID").Interface().(string)
		input["Value"] = inputReflect.FieldByName("Value").Interface()
		input["Disabled"] = inputReflect.FieldByName("Disabled").Interface().(bool)
		if val := inputReflect.FieldByName("Type"); val.IsValid() {
			input["Type"] = val.Interface().(Type)
		}
		if val := inputReflect.FieldByName("Options"); val.IsValid() {
			input["Options"] = val.Interface().([]Option)
		}
		if val := inputReflect.FieldByName("Name"); val.IsValid() {
			input["Name"] = val.Interface().(string)
		}
		if val := inputReflect.FieldByName("Prefix"); val.IsValid() {
			input["Prefix"] = val.Interface().(string)
		}
		if val := inputReflect.FieldByName("Suffix"); val.IsValid() {
			input["Suffix"] = val.Interface().(string)
		}
		input["ElementType"] = inputReflect.FieldByName("ElementType").Interface().(elementType)
		field["Input"] = input
	}
	if f.VariableToggler != nil {
		var varToggler = map[string]interface{}{}
		varTogReflect := reflect.ValueOf(f.VariableToggler)
		varToggler["AssignTo"] = varTogReflect.FieldByName("AssignTo").Interface().(string)
		varToggler["Bind"] = varTogReflect.FieldByName("Bind").Interface().(string)
		varToggler["Checked"] = varTogReflect.FieldByName("Checked").Interface().(bool)
		varToggler["Disabled"] = varTogReflect.FieldByName("Disabled").Interface().(bool)
		varToggler["ElementType"] = varTogReflect.FieldByName("ElementType").Interface().(elementType)
		field["VariableToggler"] = varToggler
	}
	return field
}

type labelElement struct {
	ElementType elementType
	Label
}

type textViewElement struct {
	ElementType elementType
	TextView
}

type inputElement struct {
	ElementType elementType
	Input
}

type textAreaElement struct {
	ElementType elementType
	TextArea
}

type selectElement struct {
	ElementType elementType
	Select
}

type variableTogglerElement struct {
	ElementType elementType
	VariableToggler
}

// Label Define a new label in a Template's Field
type Label struct {
	AssignTo string // The name of the element to assign the Label
	Text     string // The text contained in the Label
}

func (l Label) NewElement() labelElement {
	return labelElement{
		ElementType: TypeLabel,
		Label:       l,
	}
}

// TextView Define a new textview in a Template's Field
type TextView struct {
	Text       string     // The text contained in the TextView
	FontWeight FontWeight // The font-weight of the text contained in the TextView
}

// NewElement Create a new TextView in a template's Field
func (t TextView) NewElement() textViewElement {
	return textViewElement{
		ElementType: TypeTextView,
		TextView:    t,
	}
}

// Input Define a new input of varying type in a Template's Field
type Input struct {
	Bind         string      // The name of the databinder entree to bind the raw value of the Input
	BindVariable string      // The name of the databinder entree to binder the variable name of the Input
	ID           string      // The ID of Input
	Type         Type        // The type of the Input
	Value        interface{} // The value of the Input
	Disabled     bool        // Set if the input is disabled or not
	Prefix       string      // Set a Prefix for the input
	Suffix       string      // Set a Suffix for the input
}

// NewElement Create a new Input in a template's Field
func (i Input) NewElement() inputElement {
	return inputElement{
		ElementType: TypeInput,
		Input:       i,
	}
}

// TextArea Define a new textarea in a Template's Field
type TextArea struct {
	Bind         string // The name of the databinder entree to bind the raw value of the Textarea
	BindVariable string // The name of the databinder entree to binder the variable name of the Textarea
	ID           string // The ID of Textarea
	Value        string // The value of the Textarea
	Disabled     bool   // Set if the input is disabled or not
}

// NewElement Create a new TextArea input in a template's Field
func (ta TextArea) NewElement() textAreaElement {
	return textAreaElement{
		ElementType: TypeTextArea,
		TextArea:    ta,
	}
}

// Select Define a new select in a Template's Field
type Select struct {
	Bind         string      // The name of the databinder entree to bind the raw value of the Select
	BindVariable string      // The name of the databinder entree to bind the variable value of the Select
	ID           string      // The ID of Select
	Options      []Option    // The array of options contained in the Select
	Name         string      // The name of the selected Option
	Value        interface{} // The value of the selected Option
	Disabled     bool        // Set if the input is disabled or not
}

// NewElement Create a new Select input in a template's Field
func (s Select) NewElement() selectElement {
	return selectElement{
		ElementType: TypeSelect,
		Select:      s,
	}
}

// Option Define a new option in a select
type Option struct {
	ElementType elementType // Autofilled
	Name        string      // The name of the Option
	Value       interface{} // The value of the Option
}

// NewElement Create a new Option in a Select input
func (o Option) NewElement() Option {
	o.ElementType = TypeOption
	return o
}

// VariableToggler Define a new variableToggler in a Template's Field
type VariableToggler struct {
	AssignTo string // The name of the element to assign the VariableToggler
	Bind     string // The name of the databinder entree to bind the raw value of the VariableToggler
	Checked  bool   // Set the checked status of the VariableToggler, default false
	Disabled bool   // Set the disabled status of the VariableToggler, default false
}

// NewElement Create a new VariableToggler in a template's Field
func (vt VariableToggler) NewElement() variableTogglerElement {
	return variableTogglerElement{
		ElementType:     TypeCheckbox,
		VariableToggler: vt,
	}
}

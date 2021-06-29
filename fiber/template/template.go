package template

type Type string
type FontWeight string
type ElementType string

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

	TypeLabel    ElementType = "label"
	TypeTextView ElementType = "p"
	TypeInput    ElementType = "input"
	TypeTextArea ElementType = "textarea"
	TypeSelect   ElementType = "select"
	TypeOption   ElementType = "option"
	TypeCheckbox ElementType = "checkbox"

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

type Template struct {
	Fields []Field
}

// Field Define a field of a Template
type Field struct {
	Label           interface{} // The Label of the Field, use template.Label{}
	TextView        interface{} // The TextView of the Field, use template.TextView{}
	Input           interface{} // The Input of the Field, use template.Input{}, template.TextArea{}, template.Select{} or template.Checkbox{}
	VariableToggler interface{} // The VariableToggler of the Field, use template.VariableToggler{}
}

// Label Define a new label in a Template's Field
type Label struct {
	ElementType ElementType // Autofilled
	AssignTo    string      // The name of the element to assign the Label
	Text        string      // The text contained in the Label
}

func (l Label) NewElement() Label {
	l.ElementType = TypeLabel
	return l
}

// TextView Define a new textview in a Template's Field
type TextView struct {
	ElementType ElementType // Autofilled
	Text        string      // The text contained in the TextView
	FontWeight  FontWeight  // The font-weight of the text contained in the TextView
}

func (t TextView) NewElement() TextView {
	t.ElementType = TypeTextView
	return t
}

// Input Define a new input of varying type in a Template's Field
type Input struct {
	ElementType  ElementType // Autofilled
	Bind         string      // The name of the databinder entree to bind the raw value of the Input
	BindVariable string      // The name of the databinder entree to binder the variable name of the Input
	ID           string      // The ID of Input
	Type         Type        // The type of the Input
	Value        interface{} // The value of the Input
	Disabled     bool
}

func (i Input) NewElement() Input {
	i.ElementType = TypeInput
	return i
}

// TextArea Define a new textarea in a Template's Field
type TextArea struct {
	ElementType  ElementType // Autofilled
	Bind         string      // The name of the databinder entree to bind the raw value of the Textarea
	BindVariable string      // The name of the databinder entree to binder the variable name of the Textarea
	ID           string      // The ID of Textarea
	Value        string      // The value of the Textarea
}

func (ta TextArea) NewElement() TextArea {
	ta.ElementType = TypeTextArea
	return ta
}

// Select Define a new select in a Template's Field
type Select struct {
	ElementType  ElementType // Autofilled
	Bind         string      // The name of the databinder entree to bind the raw value of the Select
	BindVariable string
	ID           string      // The ID of Select
	Options      []Option    // The array of options contained in the Select
	Name         string      // The name of the selected Option
	Value        interface{} // The value of the selected Option
}

func (s Select) NewElement() Select {
	s.ElementType = TypeSelect
	return s
}

// Option Define a new option in a select
type Option struct {
	ElementType ElementType // Autofilled
	Name        string      // The name of the Option
	Value       interface{} // The value of the Option
}

func (o Option) NewElement() Option {
	o.ElementType = TypeOption
	return o
}

// VariableToggler Define a new variableToggler in a Template's Field
type VariableToggler struct {
	ElementType ElementType // Autofilled
	AssignTo    string      // The name of the element to assign the VariableToggler
	Bind        string      // The name of the databinder entree to bind the raw value of the VariableToggler
	Checked     bool        // Set the checked status of the VariableToggler, default false
	Disabled    bool        // Set the disabled status of the VariableToggler, default false
}

func (vt VariableToggler) NewElement() VariableToggler {
	vt.ElementType = TypeCheckbox
	return vt
}

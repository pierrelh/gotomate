package template

import "reflect"

type Type string

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
)

func createInput(f interface{}) map[string]interface{} {
	switch f.(type) {
	case CheckboxInput:
		return newCheckboxInput(f)
	case ColorInput:
		return newColorInput(f)
	case DateInput:
		return newDateInput(f)
	case DateTimeInput:
		return newDateTimeInput(f)
	case EmailInput:
		return newEmailInput(f)
	case FileInput:
		return newFileInput(f)
	case ImageInput:
		return newImageInput(f)
	case MonthInput:
		return newMonthInput(f)
	case NumberInput:
		return newNumberInput(f)
	case PasswordInput:
		return newPasswordInput(f)
	case Select:
		return newSelectInput(f)
	case TelInput:
		return newTelInput(f)
	case TextInput:
		return newTextInput(f)
	case TextArea:
		return newTextArea(f)
	case UrlInput:
		return newUrlInput(f)
	case WeekInput:
		return newWeekInput(f)
	}
	return nil
}

// CheckboxInput Define a new checkbox input in a Template's Field
type CheckboxInput struct {
	Bind         string // The name of the databinder entree to bind the raw value of the Input
	BindVariable string // The name of the databinder entree to binder the variable name of the Input
	Value        bool   // The value of the Input
	Disabled     bool   // Set if the input is disabled or not
}

func newCheckboxInput(f interface{}) map[string]interface{} {
	var input = map[string]interface{}{}
	reflected := reflect.ValueOf(f)
	if val := reflected.FieldByName("Bind").Interface().(string); val != "" {
		input["Bind"] = val
	}
	if val := reflected.FieldByName("BindVariable").Interface().(string); val != "" {
		input["BindVariable"] = val
	}
	if val := reflected.FieldByName("Disabled").Interface().(bool); val {
		input["Disabled"] = val
	}

	input["Value"] = reflected.FieldByName("Value").Interface().(bool)
	input["ElementType"] = TypeInput
	input["Type"] = CheckBox
	return input
}

// ColorInput Define a new color input in a Template's Field
type ColorInput struct {
	Bind         string // The name of the databinder entree to bind the raw value of the Input
	BindVariable string // The name of the databinder entree to binder the variable name of the Input
	Value        string // The value of the Input
	Disabled     bool   // Set if the input is disabled or not
}

func newColorInput(f interface{}) map[string]interface{} {
	var input = map[string]interface{}{}
	reflected := reflect.ValueOf(f)
	if val := reflected.FieldByName("Bind").Interface().(string); val != "" {
		input["Bind"] = val
	}
	if val := reflected.FieldByName("BindVariable").Interface().(string); val != "" {
		input["BindVariable"] = val
	}
	if val := reflected.FieldByName("Disabled").Interface().(bool); val {
		input["Disabled"] = val
	}

	input["Value"] = reflected.FieldByName("Value").Interface().(string)
	input["ElementType"] = TypeInput
	input["Type"] = Color
	return input
}

// DateInput Define a new date input in a Template's Field
type DateInput struct {
	Bind         string // The name of the databinder entree to bind the raw value of the Input
	BindVariable string // The name of the databinder entree to binder the variable name of the Input
	Value        string // The value of the Input
	Disabled     bool   // Set if the input is disabled or not
}

func newDateInput(f interface{}) map[string]interface{} {
	var input = map[string]interface{}{}
	reflected := reflect.ValueOf(f)
	if val := reflected.FieldByName("Bind").Interface().(string); val != "" {
		input["Bind"] = val
	}
	if val := reflected.FieldByName("BindVariable").Interface().(string); val != "" {
		input["BindVariable"] = val
	}
	if val := reflected.FieldByName("Disabled").Interface().(bool); val {
		input["Disabled"] = val
	}

	input["Value"] = reflected.FieldByName("Value").Interface().(string)
	input["ElementType"] = TypeInput
	input["Type"] = Date
	return input
}

// DateTimeInput Define a new datetime input in a Template's Field
type DateTimeInput struct {
	Bind         string // The name of the databinder entree to bind the raw value of the Input
	BindVariable string // The name of the databinder entree to binder the variable name of the Input
	Value        string // The value of the Input
	Disabled     bool   // Set if the input is disabled or not
}

func newDateTimeInput(f interface{}) map[string]interface{} {
	var input = map[string]interface{}{}
	reflected := reflect.ValueOf(f)
	if val := reflected.FieldByName("Bind").Interface().(string); val != "" {
		input["Bind"] = val
	}
	if val := reflected.FieldByName("BindVariable").Interface().(string); val != "" {
		input["BindVariable"] = val
	}
	if val := reflected.FieldByName("Disabled").Interface().(bool); val {
		input["Disabled"] = val
	}

	input["Value"] = reflected.FieldByName("Value").Interface().(string)
	input["ElementType"] = TypeInput
	input["Type"] = DatetimeLocal
	return input
}

// EmailInput Define a new email input in a Template's Field
type EmailInput struct {
	Bind         string // The name of the databinder entree to bind the raw value of the Input
	BindVariable string // The name of the databinder entree to binder the variable name of the Input
	Value        string // The value of the Input
	Disabled     bool   // Set if the input is disabled or not
}

func newEmailInput(f interface{}) map[string]interface{} {
	var input = map[string]interface{}{}
	reflected := reflect.ValueOf(f)
	if val := reflected.FieldByName("Bind").Interface().(string); val != "" {
		input["Bind"] = val
	}
	if val := reflected.FieldByName("BindVariable").Interface().(string); val != "" {
		input["BindVariable"] = val
	}
	if val := reflected.FieldByName("Disabled").Interface().(bool); val {
		input["Disabled"] = val
	}

	input["Value"] = reflected.FieldByName("Value").Interface().(string)
	input["ElementType"] = TypeInput
	input["Type"] = Email
	return input
}

// FileInput Define a new file input in a Template's Field
type FileInput struct {
	Bind         string // The name of the databinder entree to bind the raw value of the Input
	BindVariable string // The name of the databinder entree to binder the variable name of the Input
	Value        string // The value of the Input
	Disabled     bool   // Set if the input is disabled or not
}

func newFileInput(f interface{}) map[string]interface{} {
	var input = map[string]interface{}{}
	reflected := reflect.ValueOf(f)
	if val := reflected.FieldByName("Bind").Interface().(string); val != "" {
		input["Bind"] = val
	}
	if val := reflected.FieldByName("BindVariable").Interface().(string); val != "" {
		input["BindVariable"] = val
	}
	if val := reflected.FieldByName("Disabled").Interface().(bool); val {
		input["Disabled"] = val
	}

	input["Value"] = reflected.FieldByName("Value").Interface().(string)
	input["ElementType"] = TypeInput
	input["Type"] = File
	return input
}

// ImageInput Define a new image input in a Template's Field
type ImageInput struct {
	Bind         string // The name of the databinder entree to bind the raw value of the Input
	BindVariable string // The name of the databinder entree to binder the variable name of the Input
	Value        string // The value of the Input
	Disabled     bool   // Set if the input is disabled or not
}

func newImageInput(f interface{}) map[string]interface{} {
	var input = map[string]interface{}{}
	reflected := reflect.ValueOf(f)
	if val := reflected.FieldByName("Bind").Interface().(string); val != "" {
		input["Bind"] = val
	}
	if val := reflected.FieldByName("BindVariable").Interface().(string); val != "" {
		input["BindVariable"] = val
	}
	if val := reflected.FieldByName("Disabled").Interface().(bool); val {
		input["Disabled"] = val
	}

	input["Value"] = reflected.FieldByName("Value").Interface().(string)
	input["ElementType"] = TypeInput
	input["Type"] = Image
	return input
}

// MonthInput Define a new month input in a Template's Field
type MonthInput struct {
	Bind         string // The name of the databinder entree to bind the raw value of the Input
	BindVariable string // The name of the databinder entree to binder the variable name of the Input
	Value        string // The value of the Input
	Disabled     bool   // Set if the input is disabled or not
}

func newMonthInput(f interface{}) map[string]interface{} {
	var input = map[string]interface{}{}
	reflected := reflect.ValueOf(f)
	if val := reflected.FieldByName("Bind").Interface().(string); val != "" {
		input["Bind"] = val
	}
	if val := reflected.FieldByName("BindVariable").Interface().(string); val != "" {
		input["BindVariable"] = val
	}
	if val := reflected.FieldByName("Disabled").Interface().(bool); val {
		input["Disabled"] = val
	}

	input["Value"] = reflected.FieldByName("Value").Interface().(string)
	input["ElementType"] = TypeInput
	input["Type"] = Month
	return input
}

// NumberInput Define a new number input in a Template's Field
type NumberInput struct {
	Bind         string // The name of the databinder entree to bind the raw value of the Input
	BindVariable string // The name of the databinder entree to binder the variable name of the Input
	Value        int    // The value of the Input
	Decimals     int    // The number of decimals allowed
	Disabled     bool   // Set if the input is disabled or not
	Prefix       string // Set a Prefix for the input
	Suffix       string // Set a Suffix for the input
	MaxValue     int    // Set a max value for the input
	MinValue     int    // Set a min value for the input
}

func newNumberInput(f interface{}) map[string]interface{} {
	var input = map[string]interface{}{}
	reflected := reflect.ValueOf(f)
	if val := reflected.FieldByName("Bind").Interface().(string); val != "" {
		input["Bind"] = val
	}
	if val := reflected.FieldByName("BindVariable").Interface().(string); val != "" {
		input["BindVariable"] = val
	}
	if val := reflected.FieldByName("Prefix").Interface().(string); val != "" {
		input["Prefix"] = val
	}
	if val := reflected.FieldByName("Suffix").Interface().(string); val != "" {
		input["Suffix"] = val
	}
	if val := reflected.FieldByName("Disabled").Interface().(bool); val {
		input["Disabled"] = val
	}
	if val := reflected.FieldByName("Decimals").Interface().(int); val != 0 {
		input["Decimals"] = val
	}
	if val := reflected.FieldByName("MaxValue").Interface().(int); val != 0 {
		input["MaxValue"] = val
	}
	if val := reflected.FieldByName("MinValue").Interface().(int); val != 0 {
		input["MinValue"] = val
	}

	input["Value"] = reflected.FieldByName("Value").Interface().(int)
	input["ElementType"] = TypeInput
	input["Type"] = Number
	return input
}

// Option Define a new option in a select
type Option struct {
	Name  string      // The name of the Option
	Value interface{} // The value of the Option
}

// PasswordInput Define a new password input in a Template's Field
type PasswordInput struct {
	Bind         string // The name of the databinder entree to bind the raw value of the Input
	BindVariable string // The name of the databinder entree to binder the variable name of the Input
	Value        string // The value of the Input
	Disabled     bool   // Set if the input is disabled or not
}

func newPasswordInput(f interface{}) map[string]interface{} {
	var input = map[string]interface{}{}
	reflected := reflect.ValueOf(f)
	if val := reflected.FieldByName("Bind").Interface().(string); val != "" {
		input["Bind"] = val
	}
	if val := reflected.FieldByName("BindVariable").Interface().(string); val != "" {
		input["BindVariable"] = val
	}
	if val := reflected.FieldByName("Disabled").Interface().(bool); val {
		input["Disabled"] = val
	}

	input["Value"] = reflected.FieldByName("Value").Interface().(string)
	input["ElementType"] = TypeInput
	input["Type"] = Password
	return input
}

// Select Define a new select in a Template's Field
type Select struct {
	Bind         interface{} // The name of the databinder entree to bind the value of the selected Option
	BindName     string      // The name of the databinder entree to bind the name of the selected Option
	BindVariable string      // The name of the databinder entree to bind the variable value of the Select
	Options      []Option    // The array of options contained in the Select
	Disabled     bool        // Set if the input is disabled or not
}

func newSelectInput(f interface{}) map[string]interface{} {
	var input = map[string]interface{}{}
	reflected := reflect.ValueOf(f)
	if val := reflected.FieldByName("BindName").Interface().(string); val != "" {
		input["BindName"] = val
	}
	if val := reflected.FieldByName("Bind").Interface(); val != nil {
		input["Bind"] = val
	}
	if val := reflected.FieldByName("BindVariable").Interface().(string); val != "" {
		input["BindVariable"] = val
	}
	if val := reflected.FieldByName("Disabled").Interface().(bool); val {
		input["Disabled"] = val
	}
	if val := reflected.FieldByName("Options").Interface(); val != nil {
		input["Options"] = val.([]Option)
	}

	input["ElementType"] = TypeSelect
	return input
}

// TelInput Define a new tel input in a Template's Field
type TelInput struct {
	Bind         string // The name of the databinder entree to bind the raw value of the Input
	BindVariable string // The name of the databinder entree to binder the variable name of the Input
	Value        string // The value of the Input
	Disabled     bool   // Set if the input is disabled or not
}

func newTelInput(f interface{}) map[string]interface{} {
	var input = map[string]interface{}{}
	reflected := reflect.ValueOf(f)
	if val := reflected.FieldByName("Bind").Interface().(string); val != "" {
		input["Bind"] = val
	}
	if val := reflected.FieldByName("BindVariable").Interface().(string); val != "" {
		input["BindVariable"] = val
	}
	if val := reflected.FieldByName("Disabled").Interface().(bool); val {
		input["Disabled"] = val
	}

	input["Value"] = reflected.FieldByName("Value").Interface().(string)
	input["ElementType"] = TypeInput
	input["Type"] = Tel
	return input
}

// TextInput Define a new text input in a Template's Field
type TextInput struct {
	Bind         string // The name of the databinder entree to bind the raw value of the Input
	BindVariable string // The name of the databinder entree to binder the variable name of the Input
	Value        string // The value of the Input
	Disabled     bool   // Set if the input is disabled or not
	Prefix       string // Set a Prefix for the input
	Suffix       string // Set a Suffix for the input
	MaxLength    int    // Set the max length of the input
	MinLength    int    // Set the min length of the input
}

func newTextInput(f interface{}) map[string]interface{} {
	var input = map[string]interface{}{}
	reflected := reflect.ValueOf(f)
	if val := reflected.FieldByName("Bind").Interface().(string); val != "" {
		input["Bind"] = val
	}
	if val := reflected.FieldByName("BindVariable").Interface().(string); val != "" {
		input["BindVariable"] = val
	}
	if val := reflected.FieldByName("Disabled").Interface().(bool); val {
		input["Disabled"] = val
	}
	if val := reflected.FieldByName("Prefix").Interface().(string); val != "" {
		input["Prefix"] = val
	}
	if val := reflected.FieldByName("Suffix").Interface().(string); val != "" {
		input["Suffix"] = val
	}
	if val := reflected.FieldByName("MaxLength").Interface().(int); val != 0 {
		input["MaxLength"] = val
	}
	if val := reflected.FieldByName("MinLength").Interface().(int); val != 0 {
		input["MinLength"] = val
	}

	input["Value"] = reflected.FieldByName("Value").Interface().(string)
	input["ElementType"] = TypeInput
	input["Type"] = Text
	return input
}

// TextArea Define a new textarea in a Template's Field
type TextArea struct {
	Bind         string // The name of the databinder entree to bind the raw value of the Textarea
	BindVariable string // The name of the databinder entree to binder the variable name of the Textarea
	Value        string // The value of the Textarea
	Disabled     bool   // Set if the input is disabled or not
}

func newTextArea(f interface{}) map[string]interface{} {
	var input = map[string]interface{}{}
	reflected := reflect.ValueOf(f)
	if val := reflected.FieldByName("Bind").Interface().(string); val != "" {
		input["Bind"] = val
	}
	if val := reflected.FieldByName("BindVariable").Interface().(string); val != "" {
		input["BindVariable"] = val
	}
	if val := reflected.FieldByName("Disabled").Interface().(bool); val {
		input["Disabled"] = val
	}

	input["Value"] = reflected.FieldByName("Value").Interface().(string)
	input["ElementType"] = TypeTextArea
	return input
}

// UrlInput Define a new url input in a Template's Field
type UrlInput struct {
	Bind         string // The name of the databinder entree to bind the raw value of the Input
	BindVariable string // The name of the databinder entree to binder the variable name of the Input
	Value        string // The value of the Input
	Disabled     bool   // Set if the input is disabled or not
}

func newUrlInput(f interface{}) map[string]interface{} {
	var input = map[string]interface{}{}
	reflected := reflect.ValueOf(f)
	if val := reflected.FieldByName("Bind").Interface().(string); val != "" {
		input["Bind"] = val
	}
	if val := reflected.FieldByName("BindVariable").Interface().(string); val != "" {
		input["BindVariable"] = val
	}
	if val := reflected.FieldByName("Disabled").Interface().(bool); val {
		input["Disabled"] = val
	}

	input["Value"] = reflected.FieldByName("Value").Interface().(string)
	input["ElementType"] = TypeInput
	input["Type"] = Url
	return input
}

// WeekInput Define a new week input in a Template's Field
type WeekInput struct {
	Bind         string // The name of the databinder entree to bind the raw value of the Input
	BindVariable string // The name of the databinder entree to binder the variable name of the Input
	Value        string // The value of the Input
	Disabled     bool   // Set if the input is disabled or not
}

func newWeekInput(f interface{}) map[string]interface{} {
	var input = map[string]interface{}{}
	reflected := reflect.ValueOf(f)
	if val := reflected.FieldByName("Bind").Interface().(string); val != "" {
		input["Bind"] = val
	}
	if val := reflected.FieldByName("BindVariable").Interface().(string); val != "" {
		input["BindVariable"] = val
	}
	if val := reflected.FieldByName("Disabled").Interface().(bool); val {
		input["Disabled"] = val
	}

	input["Value"] = reflected.FieldByName("Value").Interface().(string)
	input["ElementType"] = TypeInput
	input["Type"] = Week
	return input
}

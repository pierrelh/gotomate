package template

import "reflect"

type FontWeight string

const (
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

// TextView Define a new textview in a Template's Field
type TextView struct {
	Text       string     // The text contained in the TextView
	FontWeight FontWeight // The font-weight of the text contained in the TextView
}

func newTextView(f interface{}) map[string]interface{} {
	var textView = map[string]interface{}{}
	tvreflect := reflect.ValueOf(f)
	if val := tvreflect.FieldByName("FontWeight").Interface().(string); val != "" {
		textView["FontWeight"] = val
	}
	if val := tvreflect.FieldByName("Text").Interface().(string); val != "" {
		textView["Text"] = val
	}
	textView["ElementType"] = TypeTextView
	return textView
}

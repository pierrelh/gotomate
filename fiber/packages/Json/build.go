package json

import (
	"fmt"

	"github.com/lxn/walk/declarative"
)

// Build Return the right databinder & the right template for a flow instruction
func Build(function string) (interface{}, []declarative.Widget) {
	switch function {
	case "CreateJson":
		return new(CreateJsonDatabinder), CreateJsonTemplate
	case "JsonToDictionary":
		return new(JsonToDictionaryDatabinder), JsonToDictionaryTemplate
	case "SaveJson":
		return new(SaveJsonDatabinder), SaveJsonTemplate
	}
	fmt.Println("GOTOMATE ERROR: Unable to find the function for instruction building")
	return nil, nil
}

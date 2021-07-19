package dictionary

import (
	"github.com/lxn/walk/declarative"
)

// Build Return the right databinder & the right template for a flow instruction
func Build(function string) (interface{}, []declarative.Widget) {
	switch function {
	case "CreateDictionary":
		return new(CreateDictionaryDatabinder), CreateDictionaryTemplate
	case "CreateEntry":
		return new(CreateEntryDatabinder), CreateEntryTemplate
	case "DictionaryToJson":
		return new(DictionaryToJsonDatabinder), DictionaryToJsonTemplate
	case "RemoveEntry":
		return new(RemoveEntryDatabinder), RemoveEntryTemplate
	}
	return nil, nil
}

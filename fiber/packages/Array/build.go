package array

import "gotomate-astilectron/fiber/template"

// Build Return the right databinder & the right template for a flow instruction
func Build(function string) (interface{}, *template.InstructionTemplate) {
	switch function {
	case "GetArrayLength":
		return new(GetArrayLengthDatabinder), GetArrayLengthTemplate
	case "GetValue":
		return new(GetValueDatabinder), GetValueTemplate
	case "PopAt":
		return new(PopAtDatabinder), PopAtTemplate
	case "PopLast":
		return new(PopLastDatabinder), PopLastTemplate
	case "PushAt":
		return new(PushAtDatabinder), PushAtTemplate
	case "PushLast":
		return new(PushLastDatabinder), PushLastTemplate
	case "RemoveAt":
		return new(RemoveAtDatabinder), RemoveAtTemplate
	case "RemoveLast":
		return new(RemoveLastDatabinder), RemoveLastTemplate
	case "Shuffle":
		return new(ShuffleDatabinder), ShuffleTemplate
	case "UpdateValue":
		return new(UpdateValueDatabinder), UpdateValueTemplate
	default:
		return nil, nil
	}
}

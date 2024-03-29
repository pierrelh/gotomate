package sleep

import (
	"gotomate/fiber/template"
)

// Build Return the right databinder & the right template for a sleep instruction
func Build(function string) (interface{}, *template.InstructionTemplate) {
	switch function {
	case "Sleep":
		return new(Databinder), SleepTemplate
	case "MilliSleep":
		return new(Databinder), SleepTemplate
	}
	return nil, nil
}

package systime

import "gotomate-astilectron/fiber/template"

// Build Return the right databinder & the right template for a systime instruction
func Build(function string) (interface{}, *template.InstructionTemplate) {
	switch function {
	case "GetCurrentSysClock":
		return new(ClockDatabinder), SysClockTemplate
	case "GetCurrentSysTime":
		return new(TimeDatabinder), SysTimeTemplate
	}
	return nil, nil
}

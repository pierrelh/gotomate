package systime

import (
	"fmt"

	"github.com/lxn/walk/declarative"
)

// Build Return the right databinder & the right template for a systime instruction
func Build(function string) (interface{}, []declarative.Widget) {
	switch function {
	case "GetCurrentSysClock":
		return new(ClockDatabinder), SysClockTemplate
	case "GetCurrentSysTime":
		return new(TimeDatabinder), SysTimeTemplate
	}
	fmt.Println("GOTOMATE ERROR: Unable to find the function for instruction building")
	return nil, nil
}

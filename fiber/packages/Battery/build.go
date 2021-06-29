package battery

import (
	"fmt"

	"github.com/lxn/walk/declarative"
)

// Build Return the right databinder & the right template for a battery instruction
func Build(function string) (interface{}, []declarative.Widget) {
	switch function {
	case "GetBattery":
		return new(UserBatDatabinder), UserBatteryTemplate
	case "GetBatteryState":
		return new(BatParameterDatabinder), ParametersTemplate
	case "GetBatteryPercentage":
		return new(BatParameterDatabinder), ParametersTemplate
	case "GetBatteryRemainingTime":
		return new(BatParameterDatabinder), ParametersTemplate
	case "GetBatteryChargeRate":
		return new(BatParameterDatabinder), ParametersTemplate
	case "GetBatteryCurrentCapacity":
		return new(BatParameterDatabinder), ParametersTemplate
	case "GetBatteryLastFullCapacity":
		return new(BatParameterDatabinder), ParametersTemplate
	case "GetBatteryDesignCapacity":
		return new(BatParameterDatabinder), ParametersTemplate
	case "GetBatteryVoltage":
		return new(BatParameterDatabinder), ParametersTemplate
	case "GetBatteryDesignVoltage":
		return new(BatParameterDatabinder), ParametersTemplate
	}
	fmt.Println("GOTOMATE ERROR: Unable to find the function for instruction building")
	return nil, nil
}

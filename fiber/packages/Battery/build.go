package battery

import "gotomate-astilectron/fiber/template"

// Build Return the right databinder & the right template for a battery instruction
func Build(function string) (interface{}, *template.InstructionTemplate) {
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
	return nil, nil
}

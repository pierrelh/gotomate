package battery

import (
	"gotomate-astilectron/fiber/variable"
	"gotomate-astilectron/log"
	"time"

	"github.com/distatus/battery"
)

// GetBattery get the first system battery if exist
func GetBattery(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Getting a battery")

	batteries, err := battery.GetAll()
	if _, isFatal := err.(battery.ErrFatal); !isFatal {
		if len(batteries) != 0 {
			errs, partialErrs := err.(battery.Errors)
			for i, bat := range batteries {
				if partialErrs && errs[i] != nil {
					log.FiberError("Error getting info for BAT%d: %s\n", i, errs[i])
					continue
				}
				variable.SetVariable(instructionData, "Output", bat)
				finished <- true
				return -1
			}
		}
	}
	finished <- true
	return -1
}

// GetBatteryChargeRate : Return a battery charge rate mW
func GetBatteryChargeRate(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Getting a battery charge rate")

	batName, err := variable.Keys{VarName: "BatteryName"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	bat := batName.(*battery.Battery)
	variable.SetVariable(instructionData, "Output", time.Duration(bat.ChargeRate))
	finished <- true
	return -1
}

// GetBatteryCurrentCapacity : Return a battery current capacity
func GetBatteryCurrentCapacity(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Getting a battery current capacity")

	batName, err := variable.Keys{VarName: "BatteryName"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	bat := batName.(*battery.Battery)
	variable.SetVariable(instructionData, "Output", bat.Current)
	finished <- true
	return -1
}

// GetBatteryDesignCapacity : Return a battery design capacity
func GetBatteryDesignCapacity(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Getting a battery design capacity")

	batName, err := variable.Keys{VarName: "BatteryName"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	bat := batName.(*battery.Battery)
	variable.SetVariable(instructionData, "Output", bat.Design)
	finished <- true
	return -1
}

// GetBatteryDesignVoltage : Return a battery design voltage
func GetBatteryDesignVoltage(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Getting a battery design voltage")

	batName, err := variable.Keys{VarName: "BatteryName"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	bat := batName.(*battery.Battery)
	variable.SetVariable(instructionData, "Output", bat.DesignVoltage)
	finished <- true
	return -1
}

// GetBatteryLastFullCapacity : Return a battery last full capacity
func GetBatteryLastFullCapacity(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Getting a battery last full capacity")

	batName, err := variable.Keys{VarName: "BatteryName"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	bat := batName.(*battery.Battery)
	variable.SetVariable(instructionData, "Output", bat.Full)
	finished <- true
	return -1
}

// GetBatteryPercentage : Return the left percentage of a battery
func GetBatteryPercentage(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Getting a battery percentage")

	batName, err := variable.Keys{VarName: "BatteryName"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	bat := batName.(*battery.Battery)
	variable.SetVariable(instructionData, "Output", bat.Current/bat.Full*100)
	finished <- true
	return -1
}

// GetBatteryRemainingTime : Return the remaining time of battery or for battery charging
func GetBatteryRemainingTime(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Getting a battery remaining time")
	var timeNum float64

	batName, err := variable.Keys{VarName: "BatteryName"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	bat := batName.(*battery.Battery)
	switch bat.State {
	case battery.Discharging:
		timeNum = bat.Current / bat.ChargeRate
	case battery.Charging:
		timeNum = (bat.Full - bat.Current) / bat.ChargeRate
	default:
		timeNum = 0
	}
	duration, _ := time.ParseDuration(log.Sprint("%fh", timeNum))
	variable.SetVariable(instructionData, "Output", duration)

	finished <- true
	return -1
}

// GetBatteryState : Return the battery state of a battery
func GetBatteryState(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Getting a battery state")

	batName, err := variable.Keys{VarName: "BatteryName"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	bat := batName.(*battery.Battery)
	variable.SetVariable(instructionData, "Output", bat.State)
	finished <- true
	return -1
}

// GetBatteryVoltage : Return a battery voltage
func GetBatteryVoltage(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Getting a battery voltage")

	batName, err := variable.Keys{VarName: "BatteryName"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	bat := batName.(*battery.Battery)
	variable.SetVariable(instructionData, "Output", bat.Voltage)
	finished <- true
	return -1
}

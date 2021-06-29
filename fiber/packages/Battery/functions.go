package battery

import (
	"fmt"
	"gotomate/fiber/variable"
	"os"
	"reflect"
	"time"

	"github.com/distatus/battery"
)

// GetBattery get the first system battery if exist
func GetBattery(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Getting battery ...")
	batteries, err := battery.GetAll()
	if _, isFatal := err.(battery.ErrFatal); !isFatal {
		if len(batteries) != 0 {
			errs, partialErrs := err.(battery.Errors)
			for i, bat := range batteries {
				if partialErrs && errs[i] != nil {
					fmt.Fprintf(os.Stderr, "FIBER ERROR: Error getting info for BAT%d: %s\n", i, errs[i])
					continue
				}
				variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), bat)
				finished <- true
				return -1
			}
		}
	}
	finished <- true
	return -1
}

// GetBatteryChargeRate : Return a battery charge rate mW
func GetBatteryChargeRate(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Getting battery charge rate ...")

	batName, err := variable.GetValue(instructionData, "BatteryName")
	if err != nil {
		finished <- true
		return -1
	}

	bat := batName.(*battery.Battery)
	variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), time.Duration(bat.ChargeRate))
	finished <- true
	return -1
}

// GetBatteryCurrentCapacity : Return a battery current capacity
func GetBatteryCurrentCapacity(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Getting battery current capacity ...")

	batName, err := variable.GetValue(instructionData, "BatteryName")
	if err != nil {
		finished <- true
		return -1
	}

	bat := batName.(*battery.Battery)
	variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), bat.Current)
	finished <- true
	return -1
}

// GetBatteryDesignCapacity : Return a battery design capacity
func GetBatteryDesignCapacity(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Getting battery design capacity ...")

	batName, err := variable.GetValue(instructionData, "BatteryName")
	if err != nil {
		finished <- true
		return -1
	}

	bat := batName.(*battery.Battery)
	variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), bat.Design)
	finished <- true
	return -1
}

// GetBatteryDesignVoltage : Return a battery design voltage
func GetBatteryDesignVoltage(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Getting battery design voltage ...")

	batName, err := variable.GetValue(instructionData, "BatteryName")
	if err != nil {
		finished <- true
		return -1
	}

	bat := batName.(*battery.Battery)
	variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), bat.DesignVoltage)
	finished <- true
	return -1
}

// GetBatteryLastFullCapacity : Return a battery last full capacity
func GetBatteryLastFullCapacity(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Getting battery last full capacity ...")

	batName, err := variable.GetValue(instructionData, "BatteryName")
	if err != nil {
		finished <- true
		return -1
	}

	bat := batName.(*battery.Battery)
	variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), bat.Full)
	finished <- true
	return -1
}

// GetBatteryPercentage : Return the left percentage of a battery
func GetBatteryPercentage(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Getting battery percentage ...")

	batName, err := variable.GetValue(instructionData, "BatteryName")
	if err != nil {
		finished <- true
		return -1
	}

	bat := batName.(*battery.Battery)
	variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), bat.Current/bat.Full*100)
	finished <- true
	return -1
}

// GetBatteryRemainingTime : Return the remaining time of battery or for battery charging
func GetBatteryRemainingTime(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Getting battery remaining time ...")
	var timeNum float64

	batName, err := variable.GetValue(instructionData, "BatteryName")
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
	duration, _ := time.ParseDuration(fmt.Sprintf("%fh", timeNum))
	variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), duration)

	finished <- true
	return -1
}

// GetBatteryState : Return the battery state of a battery
func GetBatteryState(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Getting battery state ...")

	batName, err := variable.GetValue(instructionData, "BatteryName")
	if err != nil {
		finished <- true
		return -1
	}

	bat := batName.(*battery.Battery)
	variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), bat.State)
	finished <- true
	return -1
}

// GetBatteryVoltage : Return a battery voltage
func GetBatteryVoltage(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Getting battery voltage ...")

	batName, err := variable.GetValue(instructionData, "BatteryName")
	if err != nil {
		finished <- true
		return -1
	}

	bat := batName.(*battery.Battery)
	variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), bat.Voltage)
	finished <- true
	return -1
}

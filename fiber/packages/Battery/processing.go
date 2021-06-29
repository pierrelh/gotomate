package battery

import (
	"fmt"
	"reflect"
)

// Processing process the functions from battery's package
func Processing(funcName string, instructionData reflect.Value, finished chan bool) int {
	nextID := -1
	switch funcName {
	case "GetBattery":
		go func() {
			nextID = GetBattery(instructionData, finished)
		}()
		<-finished
	case "GetBatteryChargeRate":
		go func() {
			nextID = GetBatteryChargeRate(instructionData, finished)
		}()
		<-finished
	case "GetBatteryCurrentCapacity":
		go func() {
			nextID = GetBatteryCurrentCapacity(instructionData, finished)
		}()
		<-finished
	case "GetBatteryDesignCapacity":
		go func() {
			nextID = GetBatteryDesignCapacity(instructionData, finished)
		}()
		<-finished
	case "GetBatteryDesignVoltage":
		go func() {
			nextID = GetBatteryDesignVoltage(instructionData, finished)
		}()
		<-finished
	case "GetBatteryLastFullCapacity":
		go func() {
			nextID = GetBatteryLastFullCapacity(instructionData, finished)
		}()
		<-finished
	case "GetBatteryPercentage":
		go func() {
			nextID = GetBatteryPercentage(instructionData, finished)
		}()
		<-finished
	case "GetBatteryRemainingTime":
		go func() {
			nextID = GetBatteryRemainingTime(instructionData, finished)
		}()
		<-finished
	case "GetBatteryState":
		go func() {
			nextID = GetBatteryState(instructionData, finished)
		}()
		<-finished
	case "GetBatteryVoltage":
		go func() {
			nextID = GetBatteryVoltage(instructionData, finished)
		}()
		<-finished
	default:
		fmt.Println("FIBER ERROR: This function is not integrated yet: " + funcName)
	}
	return nextID
}

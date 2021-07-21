package systime

import (
	"gotomate-astilectron/fiber/variable"
	"gotomate-astilectron/log"
	"time"
)

// GetCurrentSysClock return the current sys clock
func GetCurrentSysClock(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Getting system's clock")

	h, m, s := time.Now().Clock()
	variable.SetVariable(instructionData, "HoursOutput", h)
	variable.SetVariable(instructionData, "MinutesOutput", m)
	variable.SetVariable(instructionData, "SecondsOutput", s)
	finished <- true
	return -1
}

// GetCurrentSysTime return the current sys time
func GetCurrentSysTime(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Getting system's time")

	variable.SetVariable(instructionData, "Output", time.Now())
	finished <- true
	return -1
}

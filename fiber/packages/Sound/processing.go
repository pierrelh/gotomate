package sound

import (
	"reflect"
)

// Processing process the functions from clipboard's package
func Processing(funcName string, instructionData reflect.Value, finished chan bool) int {
	nextID := -1
	switch funcName {
	case "GetMuted":
		go func() {
			nextID = GetMuted(instructionData, finished)
		}()
		<-finished
	case "GetVolume":
		go func() {
			nextID = GetVolume(instructionData, finished)
		}()
		<-finished
	case "Mute":
		go func() {
			nextID = Mute(instructionData, finished)
		}()
		<-finished
	case "SetVolume":
		go func() {
			nextID = SetVolume(instructionData, finished)
		}()
		<-finished
	case "UnMute":
		go func() {
			nextID = UnMute(instructionData, finished)
		}()
		<-finished
	default:
		return -2
	}
	return nextID
}

package sound

import (
	"fmt"
	"gotomate/fiber/variable"
	"reflect"

	"github.com/itchyny/volume-go"
)

// GetMuted Get the status of mute
func GetMuted(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Getting Mute status ...")

	mute, _ := volume.GetMuted()
	variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), mute)
	finished <- true
	return -1
}

// GetVolume Get the current level of the volume
func GetVolume(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Get Volume ...")

	volume, _ := volume.GetVolume()
	variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), volume)
	finished <- true
	return -1
}

// Mute the system's sound
func Mute(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Muting ...")
	volume.Mute()
	finished <- true
	return -1
}

// SetVolume Set the system's volume to the wanted value
func SetVolume(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Setting Volume ...")

	volumeLevel, err := variable.GetValue(instructionData, "VolumeVarName", "VolumeIsVar", "Volume")
	if err != nil {
		finished <- true
		return -1
	}

	volume.SetVolume(volumeLevel.(int))
	finished <- true
	return -1
}

// UnMute UnMute the system's sound
func UnMute(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: UnMuting ...")
	volume.Unmute()
	finished <- true
	return -1
}

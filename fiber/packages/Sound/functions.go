package sound

import (
	"gotomate-astilectron/fiber/variable"
	"gotomate-astilectron/log"
	"reflect"

	"github.com/itchyny/volume-go"
)

// GetMuted Get the status of mute
func GetMuted(instructionData reflect.Value, finished chan bool) int {
	log.FiberInfo("Getting mute's status")

	mute, _ := volume.GetMuted()
	variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), mute)
	finished <- true
	return -1
}

// GetVolume Get the current level of the volume
func GetVolume(instructionData reflect.Value, finished chan bool) int {
	log.FiberInfo("Getting volume's level")

	volume, _ := volume.GetVolume()
	variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), volume)
	finished <- true
	return -1
}

// Mute the system's sound
func Mute(instructionData reflect.Value, finished chan bool) int {
	log.FiberInfo("Muting system's volume")
	volume.Mute()
	finished <- true
	return -1
}

// SetVolume Set the system's volume to the wanted value
func SetVolume(instructionData reflect.Value, finished chan bool) int {
	log.FiberInfo("Setting system's volume to wanted value")

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
	log.FiberInfo("Unmuting system's volume")
	volume.Unmute()
	finished <- true
	return -1
}

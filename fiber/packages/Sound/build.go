package sound

import (
	"fmt"

	"github.com/lxn/walk/declarative"
)

// Build Return the right databinder & the right template for a flow instruction
func Build(function string) (interface{}, []declarative.Widget) {
	switch function {
	case "GetMuted":
		return new(GetMutedDatabinder), GetMutedTemplate
	case "GetVolume":
		return new(GetVolumeDatabinder), GetVolumeTemplate
	case "Mute":
		return new(MuteVolumeDatabinder), nil
	case "SetVolume":
		return new(SetVolumeDatabinder), SetVolumeTemplate
	case "UnMute":
		return new(UnMuteVolumeDatabinder), nil
	default:
		fmt.Println("GOTOMATE ERROR: Unable to find the function for instruction building")
		return nil, nil
	}
}

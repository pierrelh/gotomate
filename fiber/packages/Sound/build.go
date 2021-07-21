package sound

import "gotomate-astilectron/fiber/template"

// Build Return the right databinder & the right template for a flow instruction
func Build(function string) (interface{}, *template.InstructionTemplate) {
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
		return nil, nil
	}
}

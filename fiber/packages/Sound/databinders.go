package sound

// GetMutedDatabinder Define the GetMuted parameters
type GetMutedDatabinder struct {
	Output string
}

// GetVolumeDatabinder Define the GetVolume parameters
type GetVolumeDatabinder struct {
	Output string
}

// MuteVolumeDatabinder Define the MuteVolume parameters
type MuteVolumeDatabinder struct {
}

// SetVolumeDatabinder Define the SetVolume parameters
type SetVolumeDatabinder struct {
	Volume        int
	VolumeVarName string
	VolumeIsVar   bool
}

// UnMuteVolumeDatabinder Define the UnMuteVolume parameters
type UnMuteVolumeDatabinder struct {
}

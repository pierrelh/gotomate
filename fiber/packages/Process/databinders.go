package process

// GetTitleDatabinder Define the GetTitle databinder
type GetTitleDatabinder struct {
	PID        int
	PIDVarName string
	PIDIsVar   bool
	Output     string
}

// GetPidDatabinder Define the GetPid databinder
type GetPidDatabinder struct {
	Output string
}

// KillProcessDatabinder Define the KillProcess databinder
type KillProcessDatabinder struct {
	PID        int
	PIDVarName string
	PIDIsVar   bool
}

// StartProcessDatabinder Define the StartProcess databinder
type StartProcessDatabinder struct {
	Path        string
	PathVarName string
	PathIsVar   bool
	Output      string
}

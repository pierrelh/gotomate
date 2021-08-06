package command

// DirDatabinder Define the Dir parameters
type DirDatabinder struct {
	CommandVarName string
	Dir            string
	DirIsVar       bool
	DirVarName     string
}

// GetOutputDatabinder Define the GetOutput parameters
type GetOutputDatabinder struct {
	CommandVarName string
	Output         string
}

// NewDatabinder Define the New parameters
type NewDatabinder struct {
	Program          string
	ProgramIsVar     bool
	ProgramVarName   string
	Parameter        string
	ParameterIsVar   bool
	ParameterVarName string
	Output           string
}

// RunDatabinder Define the Run parameters
type RunDatabinder struct {
	CommandVarName string
}

// StartDatabinder Define the Start parameters
type StartDatabinder struct {
	CommandVarName string
}

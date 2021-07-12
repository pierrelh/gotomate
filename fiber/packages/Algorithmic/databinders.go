package algorithmic

// ForDatabinder Define the For parameters
type ForDatabinder struct {
	ValueOneVarName         string
	Comparator              string
	ValueTwo                int
	ValueTwoVarName         string
	ValueTwoIsVar           bool
	IncrementVarName        string
	Increment               int
	IncrementIsVar          bool
	FalseInstruction        int
	FalseInstructionVarName string
	FalseInstructionIsVar   bool
}

// IfDatabinder Define the If parameters
type IfDatabinder struct {
	ValueOne                interface{}
	ValueOneVarName         string
	ValueOneIsVar           bool
	Comparator              string
	ValueTwo                interface{}
	ValueTwoVarName         string
	ValueTwoIsVar           bool
	FalseInstruction        int
	FalseInstructionVarName string
	FalseInstructionIsVar   bool
}

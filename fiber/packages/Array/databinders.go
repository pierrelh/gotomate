package array

// GetArrayLengthDatabinder Define the GetArrayLength parameters
type GetArrayLengthDatabinder struct {
	ArrayVarName string
	Output       string
}

// GetValueDatabinder Define the GetValue parameters
type GetValueDatabinder struct {
	ArrayVarName string
	Index        int
	IndexVarName string
	IndexIsVar   bool
	Output       string
}

// PopAtDatabinder Define the PopAt parameters
type PopAtDatabinder struct {
	ArrayVarName string
	Index        int
	IndexVarName string
	IndexIsVar   bool
	Output       string
}

// PopLastDatabinder Define the PopLast parameters
type PopLastDatabinder struct {
	ArrayVarName string
	Output       string
}

// PushAtDatabinder Define the PushAt parameters
type PushAtDatabinder struct {
	ArrayVarName string
	Index        int
	IndexVarName string
	IndexIsVar   bool
	ValueVarName string
}

// PushLastDatabinder Define the PushLast parameters
type PushLastDatabinder struct {
	ArrayVarName string
	ValueVarName string
}

// RemoveAtDatabinder Define the RemoveAt parameters
type RemoveAtDatabinder struct {
	ArrayVarName string
	Index        int
	IndexVarName string
	IndexIsVar   bool
}

// RemoveLastDatabinder Define the RemoveLast parameters
type RemoveLastDatabinder struct {
	ArrayVarName string
}

// ShuffleDatabinder Define the Shuffle parameters
type ShuffleDatabinder struct {
	ArrayVarName string
}

// UpdateValueDatabinder Define the UpdateValue parameters
type UpdateValueDatabinder struct {
	ArrayVarName string
	Index        int
	IndexVarName string
	IndexIsVar   bool
	ValueVarName string
}

package clipboard

// WriteDatabinder Define the ClipboardWrite parameters
type WriteDatabinder struct {
	Content string
	VarName string
	IsVar   bool
}

// ReadDatabinder Define the ClipboardRead parameters
type ReadDatabinder struct {
	Output string
}

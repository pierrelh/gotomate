package file

// CreateDatabinder Define the Create parameters
type CreateDatabinder struct {
	Path        string
	PathVarName string
	PathIsVar   bool
}

// DeleteDatabinder Define the  Delete parameters
type DeleteDatabinder struct {
	Path        string
	PathVarName string
	PathIsVar   bool
}

// ReadDatabinder Define the Read parameters
type ReadDatabinder struct {
	Path        string
	PathVarName string
	PathIsVar   bool
	Output      string
}

// WriteDatabinder Define the Write parameters
type WriteDatabinder struct {
	Path           string
	PathVarName    string
	PathIsVar      bool
	Content        string
	ContentVarName string
	ContentIsVar   bool
}

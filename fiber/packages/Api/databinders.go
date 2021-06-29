package api

// GetDatabinder Define the Get parameters
type GetDatabinder struct {
	Path        string
	PathVarName string
	PathIsVar   bool
	Output      string
}

// PostDatabinder Define the Post parameters
type PostDatabinder struct {
	DataVarName string
	Path        string
	PathVarName string
	PathIsVar   bool
	Output      string
}

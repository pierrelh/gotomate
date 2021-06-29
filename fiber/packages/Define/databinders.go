package define

// ArrayDatabinder Define the Array parameters
type ArrayDatabinder struct {
	Name    string
	Value   string
	VarName string
	IsVar   bool
}

// BoolDatabinder Define the Bool parameters
type BoolDatabinder struct {
	Name    string
	Value   bool
	VarName string
	IsVar   bool
}

// FloatDatabinder Define the Float parameters
type FloatDatabinder struct {
	Name    string
	Value   float64
	VarName string
	IsVar   bool
}

// IntDatabinder Define the Int parameters
type IntDatabinder struct {
	Name    string
	Value   int
	VarName string
	IsVar   bool
}

// StringDatabinder Define the String parameters
type StringDatabinder struct {
	Name    string
	Value   string
	VarName string
	IsVar   bool
}

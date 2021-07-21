package keyboard

// TapDatabinder Define the KeyboardTap parameters
type TapDatabinder struct {
	Input    string
	Special1 string
	Special2 string
}

// TypeDatabinder Define the KeyboardType parameters
type TypeDatabinder struct {
	Input      string
	VarName    string
	InputIsVar bool
}

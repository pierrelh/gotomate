package keyboard

// Inputs Define the possibles values of KeyboardInput
func Inputs() []*InputDatabinder {
	return []*InputDatabinder{
		{""},
		{"alt"},
		{"cmd"},
		{"shift"},
		{"ctrl"},
		{"enter"},
	}
}

// InputDatabinder Define the KeyboardInput parameters
type InputDatabinder struct {
	Name string
}

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

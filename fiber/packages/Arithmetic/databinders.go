package arithmetic

// ArithmeticDatabinder Define the Arithmetic parameters
type ArithmeticDatabinder struct {
	ValueOne      float64
	VarOneName    string
	ValueOneIsVar bool
	ValueTwo      float64
	VarTwoName    string
	ValueTwoIsVar bool
	Output        string
}

// SqrtDatabinder Define the Sqrt parameters
type SqrtDatabinder struct {
	Value      float64
	VarName    string
	ValueIsVar bool
	Output     string
}

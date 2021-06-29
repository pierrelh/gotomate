package conversion

// BoolConversionDatabinder Define the BoolConversion parameters
type BoolConversionDatabinder struct {
	Input        bool
	InputVarName string
	InputIsVar   bool
	Output       string
}

// FloatConversionDatabinder Define the FloatConversion parameters
type FloatConversionDatabinder struct {
	Input        float64
	InputVarName string
	InputIsVar   bool
	Output       string
}

// IntConversionDatabinder Define the IntConversion parameters
type IntConversionDatabinder struct {
	Input        int
	InputVarName string
	InputIsVar   bool
	Output       string
}

// StringConversionDatabinder Define the StringConversion parameters
type StringConversionDatabinder struct {
	Input        string
	InputVarName string
	InputIsVar   bool
	Output       string
}

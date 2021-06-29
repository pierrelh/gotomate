package screen

// MouseColorDatabinder Define the GetMouseColor databinder
type MouseColorDatabinder struct {
	Output string
}

// PixelColorDatabinder Define the GetPixelColor databinder
type PixelColorDatabinder struct {
	X        int
	XVarName string
	XIsVar   bool
	Y        int
	YVarName string
	YIsVar   bool
	Output   string
}

// SizeDatabinder Define the GetScreenSize databinder
type SizeDatabinder struct {
	HeightOutput string
	WidthOutput  string
}

// PartScreenShotDatabinder Define PartScreenShot databinder
type PartScreenShotDatabinder struct {
	Path        string
	PathVarName string
	PathIsVar   bool
	X           int
	XVarName    string
	XIsVar      bool
	Y           int
	YVarName    string
	YIsVar      bool
	W           int
	WVarName    string
	WIsVar      bool
	H           int
	HVarName    string
	HIsVar      bool
}

// ScreenShotDatabinder Define ScreenShot databinder
type ScreenShotDatabinder struct {
	Path        string
	PathVarName string
	PathIsVar   bool
}

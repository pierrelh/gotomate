package screen

import "gotomate-astilectron/fiber/template"

// Build Return the right databinder & the right template for a screen instruction
func Build(function string) (interface{}, *template.InstructionTemplate) {
	switch function {
	case "GetPixelColor":
		return new(PixelColorDatabinder), PixelColorTemplate
	case "GetMouseColor":
		return new(MouseColorDatabinder), MouseColorTemplate
	case "GetScreenSize":
		return new(SizeDatabinder), ScreenSizeTemplate
	case "PartScreenShot":
		return new(PartScreenShotDatabinder), PartScreenShotTemplate
	case "ScreenShot":
		return new(ScreenShotDatabinder), ScreenShotTemplate
	}
	return nil, nil
}

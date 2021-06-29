package screen

import (
	"fmt"

	"github.com/lxn/walk/declarative"
)

// Build Return the right databinder & the right template for a screen instruction
func Build(function string) (interface{}, []declarative.Widget) {
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
	fmt.Println("GOTOMATE ERROR: Unable to find the function for instruction building")
	return nil, nil
}

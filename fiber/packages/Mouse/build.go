package mouse

import (
	"github.com/lxn/walk/declarative"
)

// Build Return the right databinder & the right template for a mouse instruction
func Build(function string) (interface{}, []declarative.Widget) {
	switch function {
	case "Click":
		return new(ClickDatabinder), ClickTemplate
	case "Drag":
		return new(DragDatabinder), DragTemplate
	case "DragSmooth":
		return new(DragDatabinder), DragTemplate
	case "Move":
		return new(MoveDatabinder), MoveTemplate
	case "MoveSmooth":
		return new(MoveDatabinder), MoveTemplate
	case "Scroll":
		return new(ScrollDatabinder), ScrollTemplate
	}
	return nil, nil
}

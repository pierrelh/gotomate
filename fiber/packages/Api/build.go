package api

import (
	"github.com/lxn/walk/declarative"
)

// Build Return the right databinder & the right template for a flow instruction
func Build(function string) (interface{}, []declarative.Widget) {
	switch function {
	case "Get":
		return new(GetDatabinder), GetTemplate
	case "Post":
		return new(PostDatabinder), PostTemplate
	}
	return nil, nil
}

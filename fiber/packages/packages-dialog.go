package packages

import (
	"fmt"
	// DON'T REMOVE ME / New packages inserted here
	algorithmic "gotomate/fiber/packages/Algorithmic"
	api "gotomate/fiber/packages/Api"
	arithmetic "gotomate/fiber/packages/Arithmetic"
	array "gotomate/fiber/packages/Array"
	battery "gotomate/fiber/packages/Battery"
	chronometer "gotomate/fiber/packages/Chronometer"
	clipboard "gotomate/fiber/packages/Clipboard"
	conversion "gotomate/fiber/packages/Conversion"
	define "gotomate/fiber/packages/Define"
	dictionary "gotomate/fiber/packages/Dictionary"
	file "gotomate/fiber/packages/File"
	flow "gotomate/fiber/packages/Flow"
	input "gotomate/fiber/packages/Input"
	gotomatejson "gotomate/fiber/packages/Json"
	keyboard "gotomate/fiber/packages/Keyboard"
	log "gotomate/fiber/packages/Log"
	mouse "gotomate/fiber/packages/Mouse"
	notification "gotomate/fiber/packages/Notification"
	process "gotomate/fiber/packages/Process"
	scraping "gotomate/fiber/packages/Scraping"
	screen "gotomate/fiber/packages/Screen"
	sleep "gotomate/fiber/packages/Sleep"
	sound "gotomate/fiber/packages/Sound"
	systime "gotomate/fiber/packages/Systime"

	"github.com/lxn/walk"
	declarative "github.com/lxn/walk/declarative"
)

// Dialog Setting the requirement of a dialog window
type Dialog struct {
	Dialog        *walk.Dialog
	DialogContent declarative.Dialog
	DataBinder    *walk.DataBinder
	AcceptButton  *walk.PushButton
	CancelButton  *walk.PushButton
}

// CreateNewDialog Create the dialog for the function & add the needed template
func CreateNewDialog(packageName string, funcName string, databinder ...interface{}) (interface{}, *Dialog) {
	dialog := new(Dialog)
	source, children := PackageDecode(packageName, funcName)

	if len(databinder) != 0 {
		source = databinder[0]
	}

	dialog.DialogContent = declarative.Dialog{
		Icon:          "/img/icon.ico",
		Title:         funcName + " Settings",
		Font:          declarative.Font{Family: "Roboto", PointSize: 9},
		AssignTo:      &dialog.Dialog,
		DefaultButton: &dialog.AcceptButton,
		CancelButton:  &dialog.CancelButton,
		DataBinder: declarative.DataBinder{
			AssignTo:       &dialog.DataBinder,
			Name:           funcName,
			DataSource:     source,
			ErrorPresenter: declarative.ToolTipErrorPresenter{},
		},
		MinSize: declarative.Size{
			Width:  300,
			Height: 300,
		},
		Layout: declarative.VBox{},
		Children: []declarative.Widget{
			declarative.Composite{
				Layout:   declarative.VBox{},
				Children: children,
			},
			declarative.Composite{
				Alignment: declarative.Alignment2D(walk.AlignHCenterVCenter),
				Layout:    declarative.HBox{},
				Children: []declarative.Widget{
					declarative.PushButton{
						AssignTo: &dialog.AcceptButton,
						Text:     "OK",
						OnClicked: func() {
							if err := dialog.DataBinder.Submit(); err != nil {
								fmt.Println("GOTOMATE ERROR: Unable to send the dialog's datas")
								return
							}

							dialog.Dialog.Accept()
						},
					},
					declarative.PushButton{
						AssignTo:  &dialog.CancelButton,
						Text:      "Cancel",
						OnClicked: func() { dialog.Dialog.Cancel() },
					},
				},
			},
		},
	}

	return source, dialog
}

// PackageDecode Getting the right databinder & the right template needed
func PackageDecode(packageName string, funcName string) (interface{}, []declarative.Widget) {
	switch packageName {
	case "Flow":
		return flow.Build(funcName)
	// DON'T REMOVE ME / New Build inserted here
	case "Algorithmic":
		return algorithmic.Build(funcName)
	case "Api":
		return api.Build(funcName)
	case "Arithmetic":
		return arithmetic.Build(funcName)
	case "Array":
		return array.Build(funcName)
	case "Battery":
		return battery.Build(funcName)
	case "Chronometer":
		return chronometer.Build(funcName)
	case "Clipboard":
		return clipboard.Build(funcName)
	case "Conversion":
		return conversion.Build(funcName)
	case "Define":
		return define.Build(funcName)
	case "Dictionary":
		return dictionary.Build(funcName)
	case "File":
		return file.Build(funcName)
	case "Input":
		return input.Build(funcName)
	case "Json":
		return gotomatejson.Build(funcName)
	case "Keyboard":
		return keyboard.Build(funcName)
	case "Log":
		return log.Build(funcName)
	case "Mouse":
		return mouse.Build(funcName)
	case "Notification":
		return notification.Build(funcName)
	case "Process":
		return process.Build(funcName)
	case "Scraping":
		return scraping.Build(funcName)
	case "Screen":
		return screen.Build(funcName)
	case "Sleep":
		return sleep.Build(funcName)
	case "Sound":
		return sound.Build(funcName)
	case "Systime":
		return systime.Build(funcName)
	default:
		fmt.Println("GOTOMATE ERROR: Unable to find the dialog's package")
		return nil, nil
	}
}

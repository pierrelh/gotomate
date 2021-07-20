// Print a new notification in windows system

package notification

import (
	"gotomate-astilectron/fiber/variable"
	"gotomate-astilectron/log"

	"github.com/lxn/walk"
)

// Create create a new notification with presets & push it
func Create(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Creating a notification")

	title, err := variable.Keys{VarName: "TitleVarName", IsVarName: "TitleIsVar", Name: "Title"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	msg, err := variable.Keys{VarName: "MessageVarName", IsVarName: "MessageIsVar", Name: "Message"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	notTitle := "Default Title"
	notMsg := "Default Message"
	if title != "" {
		notTitle = title.(string)
	}
	if msg != "" {
		notMsg = msg.(string)
	}

	mw, err := walk.NewMainWindow()
	if err != nil {
		log.FiberError("FIBER ERROR: ", err)
	}

	icon, err := walk.Resources.Icon("/img/icon.ico")
	if err != nil {
		log.FiberError("FIBER ERROR: ", err)
	}

	ni, err := walk.NewNotifyIcon(mw)
	if err != nil {
		log.FiberError("FIBER ERROR: ", err)
	}

	if err := ni.SetIcon(icon); err != nil {
		log.FiberError("FIBER ERROR: ", err)
	}

	if err := ni.SetVisible(true); err != nil {
		log.FiberError("FIBER ERROR: ", err)
	}

	if err := ni.ShowInfo(notTitle, notMsg); err != nil {
		log.FiberError("FIBER ERROR: ", err)
	}

	finished <- true

	mw.Run()
	return -1
}

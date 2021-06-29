// Print a new notification in windows system

package notification

import (
	"fmt"
	"gotomate/fiber/variable"
	"reflect"

	"github.com/lxn/walk"
)

// Create create a new notification with presets & push it
func Create(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Creating notification ...")

	title, err := variable.GetValue(instructionData, "TitleVarName", "TitleIsVar", "Title")
	if err != nil {
		finished <- true
		return -1
	}

	msg, err := variable.GetValue(instructionData, "MessageVarName", "MessageIsVar", "Message")
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
		fmt.Println("FIBER ERROR: ", err)
	}

	icon, err := walk.Resources.Icon("/img/icon.ico")
	if err != nil {
		fmt.Println("FIBER ERROR: ", err)
	}

	ni, err := walk.NewNotifyIcon(mw)
	if err != nil {
		fmt.Println("FIBER ERROR: ", err)
	}

	if err := ni.SetIcon(icon); err != nil {
		fmt.Println("FIBER ERROR: ", err)
	}

	if err := ni.SetVisible(true); err != nil {
		fmt.Println("FIBER ERROR: ", err)
	}

	if err := ni.ShowInfo(notTitle, notMsg); err != nil {
		fmt.Println("FIBER ERROR: ", err)
	}

	finished <- true

	mw.Run()
	return -1
}

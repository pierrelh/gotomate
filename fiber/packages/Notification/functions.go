package notification

import (
	"gotomate/fiber/variable"
	"gotomate/globals"
	"gotomate/log"

	"gopkg.in/toast.v1"
)

// Create create a new notification with presets
func Create(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Creating a notification")

	title, err1 := variable.Keys{VarName: "TitleVarName", IsVarName: "TitleIsVar", Name: "Title"}.GetValue(instructionData)
	msg, err2 := variable.Keys{VarName: "MessageVarName", IsVarName: "MessageIsVar", Name: "Message"}.GetValue(instructionData)

	if err1 != nil || err2 != nil {
		finished <- true
		return -1
	}

	notification := &toast.Notification{
		AppID:   globals.AppName,
		Title:   title.(string),
		Message: msg.(string),
		Icon:    globals.DirectoryPath + globals.AppIcon,
		// Actions: []toast.Action{
		// 	{"protocol", "I'm a button", ""},
		// 	{"protocol", "Me too!", ""},
		// },
	}
	variable.SetVariable(instructionData, "Output", notification)

	finished <- true
	return -1
}

// Push an existing notification
func Push(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Pushing a notification")

	notification, err := variable.Keys{VarName: "NotificationVarName"}.GetValue(instructionData)
	if err != nil {
		finished <- true
		return -1
	}

	err = notification.(*toast.Notification).Push()
	if err != nil {
		log.FiberError("Creating a notification")
	}

	finished <- true
	return -1
}

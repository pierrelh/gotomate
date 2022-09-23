package main

import (
	"gotomate/fiber"
	"gotomate/fiber/packages"
	"gotomate/log"
	"gotomate/resources/app"
	"gotomate/resources/app/files"
	"gotomate/resources/app/message"
	"os"

	"github.com/asticode/go-astilectron"
)

var newFiber = fiber.NewFiber
var a = new(app.App).Build()

func main() {
	// Building & Creating the app
	a.Start()
	a.Create()

	// Open dev tools
	a.Window.OpenDevTools()

	// Create the menu
	a.CreateMenu()

	getEvents()

	// Initializing the Gotomate's packages
	a.Window.SendMessage(
		message.New("InitPackages", new(packages.GotomatePackages).Hydrate()),
	)

	// Initializing the first instruction & the fiber's instructions
	a.Window.SendMessage(
		message.New("NewFiber", newFiber.New()),
	)

	// Blocking pattern
	a.Asti.Wait()
}

// getEvents Receive the message sended by the gui & process them
func getEvents() {
	a.Window.OnMessage(func(m *astilectron.EventMessage) interface{} {

		// Unmarshal
		var s *message.Message
		m.Unmarshal(&s)
		content := s.Content.(map[string]interface{})

		// Process message
		switch s.Identifier {
		case "CreateInstruction":
			return newFiber.CreateInstruction(content)

		case "UpdateInstructionPosition":
			if instruction := newFiber.Instructions.FindInstructionById(int(content["ID"].(float64))); instruction != nil && content["X"] != nil && content["Y"] != nil {
				instruction.UpdatePosition(int(content["X"].(float64)), int(content["Y"].(float64)))
			}

		case "DeleteInstruction":
			if instruction := newFiber.Instructions.FindInstructionById(int(content["ID"].(float64))); instruction != nil {
				if delete := newFiber.Instructions.DeleteInstruction(instruction); delete {
					return message.New("true", instruction.ID)
				}
			}
			return nil

		case "GetInstructionTemplate":
			if instruction := newFiber.Instructions.FindInstructionById(int(content["ID"].(float64))); instruction != nil {
				data := map[string]interface{}{
					"Template":   &instruction.Template,
					"Databinder": instruction.Datas,
				}
				return message.New("true", data)
			}
			return nil

		case "SetDatabinderDatas":
			instruction := newFiber.Instructions.FindInstructionById(int(content["ID"].(float64)))
			instruction.UpdateDatabinder(content["Databinder"].(map[string]interface{}))

		case "IWCreateNewFiber":
			a.Window.SendMessage(
				message.New("NewFiber", newFiber.New()),
			)

		case "UpdateFiberName":
			newFiber.SetName(content["name"].(string))

		case "IWOpenSavedFiber":
			a.Window.SendMessage(
				message.New("NewFiber", fiber.NewFiber.Open(content["data"].(string))),
			)

		case "RunFiber":
			go fiber.NewFiber.Run()

		case "StopFiber":
			go fiber.NewFiber.Stop()

		case "UpdateInstructionNextID":
			if instruction := newFiber.Instructions.FindInstructionById(int(content["ID"].(float64))); instruction != nil {
				instruction.UpdateNextID(int(content["NextID"].(float64)))
			}

		case "GoTroughFolder":
			extension := content["Extension"].(string)
			action := content["Action"].(string)
			content, path := files.GetHomeJson(content["Path"].(string), extension)
			data := map[string]interface{}{
				"Path":      path,
				"Files":     content,
				"Extension": extension,
			}
			a.Window.SendMessage(
				message.New(action, data),
			)

		case "ImportFiber":
			a.Window.SendMessage(
				message.New("NewFiber", fiber.NewFiber.Import(content["File"].(string))),
			)

		case "IWImportFiber":
			extension := ".json"
			content, path := files.GetHomeJson("home", extension)
			data := map[string]interface{}{
				"Path":      path,
				"Files":     content,
				"Extension": extension,
			}
			a.Window.SendMessage(
				message.New("ImportFiber", data),
			)

		case "ExportFiber":
			fiber.NewFiber.Export(content["File"].(string))

		case "ImportPackage":
			err := packages.ImportPackage(content["File"].(string))
			if err != nil {
				log.GotomateError(err)
			}
			os.Exit(1)

		default:
			log.GotomateError("Unknown identifier received: ", s.Identifier)
			return nil

		}
		return nil
	})
}

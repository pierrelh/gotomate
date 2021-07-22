package fiber

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	// DON'T REMOVE ME / New packages inserted here

	"gotomate/fiber/instructions"
	"gotomate/fiber/packages"
	algorithmicpack "gotomate/fiber/packages/Algorithmic"
	apipack "gotomate/fiber/packages/Api"
	arithmeticpack "gotomate/fiber/packages/Arithmetic"
	arraypack "gotomate/fiber/packages/Array"
	batterypack "gotomate/fiber/packages/Battery"
	chronometerpack "gotomate/fiber/packages/Chronometer"
	clipboardpack "gotomate/fiber/packages/Clipboard"
	conversionpack "gotomate/fiber/packages/Conversion"
	definepack "gotomate/fiber/packages/Define"
	dictionarypack "gotomate/fiber/packages/Dictionary"
	filepack "gotomate/fiber/packages/File"
	flowpack "gotomate/fiber/packages/Flow"
	inputpack "gotomate/fiber/packages/Input"
	jsonpack "gotomate/fiber/packages/Json"
	keyboardpack "gotomate/fiber/packages/Keyboard"
	logpack "gotomate/fiber/packages/Log"
	mousepack "gotomate/fiber/packages/Mouse"
	notificationpack "gotomate/fiber/packages/Notification"
	processpack "gotomate/fiber/packages/Process"
	scrapingpack "gotomate/fiber/packages/Scraping"
	screenpack "gotomate/fiber/packages/Screen"
	sleeppack "gotomate/fiber/packages/Sleep"
	soundpack "gotomate/fiber/packages/Sound"
	systimepack "gotomate/fiber/packages/Systime"
	"gotomate/fiber/variable"
	"gotomate/log"
	"reflect"
	"sort"
)

// NewFiber Define the new automate's fiber
var NewFiber = new(Fiber)
var running = 0
var finished = make(chan bool)
var stop = make(chan bool)

const (
	finishedFiber   int = -99
	unknownFunction int = -2
	continuFiber    int = -1
)

// LoadingFiber Initialize the loading fiber structure
type LoadingFiber struct {
	Name         string
	Instructions instructions.FiberLoadingInstruction
}

// Fiber Initialize the fiber structure
type Fiber struct {
	Name         string
	Instructions instructions.FiberInstructions
}

// New Clean the current fiber a hydrate herself with a new instructions that she return
func (fiber *Fiber) New() *Fiber {
	fiber.Clean()
	inst := new(instructions.FiberInstructions).Hydrate()
	fiber.Instructions = append(fiber.Instructions, inst)
	return fiber
}

// SetName set the name of the fiber
func (fiber *Fiber) SetName(name string) {
	fiber.Name = name
}

// CreateInstruction Create a instruction in the fiber
func (fiber *Fiber) CreateInstruction(content map[string]interface{}) *instructions.Instruction {
	id := 1
	if len(fiber.Instructions) > 1 {
		id = fiber.Instructions[len(fiber.Instructions)-1].ID + 1
	}

	inst := new(instructions.Instruction).Build(id, content)
	fiber.Instructions = append(fiber.Instructions, inst)
	return inst
}

// Save Save the current fiber
func (fiber *Fiber) Save(filePath ...string) {

	// Getting the path where to save the fiber
	path := "saves/"
	if len(filePath) > 0 {
		path = filePath[0] + "/"
	}

	toSave := fiber

	// Removing the templates from the save
	for i := 0; i < len(toSave.Instructions); i++ {
		toSave.Instructions[i].Template = nil
	}

	// Saving the fiber
	fullPath := path + toSave.Name + ".json"
	file, _ := json.Marshal(fiber)
	ioutil.WriteFile(fullPath, file, 0644)
}

// IsSaved Check if the fiber is saved
func (fiber *Fiber) IsSaved() bool {
	saved := false
	file, _ := json.Marshal(fiber)
	root := "./saves"

	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".json" {
			fullPath, _ := filepath.Abs(path)
			jsonFile, _ := os.Open(fullPath)
			byteValue, _ := ioutil.ReadAll(jsonFile)
			if bytes.Equal(file, byteValue) {
				saved = true
			}
		}
		return nil
	})

	return saved
}

// Open Open a saved fiber from the menu My Fibers
func (fiber *Fiber) Open(path string) *Fiber {
	jsonFile, err := os.Open(path)

	if err != nil {
		return nil
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var loadingFiber LoadingFiber
	err = json.Unmarshal(byteValue, &loadingFiber)
	if err != nil {
		log.GotomateError("Unable to extract the saved fiber")
	}

	fiber.Clean()
	fiber.Name = loadingFiber.Name

	for _, instruction := range loadingFiber.Instructions {
		databinder, fields := packages.PackageDecode(instruction.Package, instruction.FuncName)

		newInstruction := &instructions.Instruction{
			ID:       instruction.ID,
			Package:  instruction.Package,
			FuncName: instruction.FuncName,
			X:        instruction.X,
			Y:        instruction.Y,
			IconPath: instruction.IconPath,
			NextID:   instruction.NextID,
			Template: fields,
		}
		if databinder != nil {
			if err := json.Unmarshal(instruction.Datas, databinder); err != nil {
				log.GotomateError("Unable to convert the saved instruction")
			}
			newInstruction.Datas = databinder
		}

		fiber.Instructions = append(fiber.Instructions, newInstruction)
	}
	return fiber
}

// Export Export the current fiber to a specified path
func (fiber *Fiber) Export(filePath string) {
	fiber.Save(filePath)
}

// Import Import a new fiber from a specified path
func (fiber *Fiber) Import(filePath string) *Fiber {
	return fiber.Open(filePath)
}

// Clean Delete all the instructions of the current fiber
func (fiber *Fiber) Clean() {
	p := reflect.ValueOf(fiber).Elem()
	p.Set(reflect.Zero(p.Type()))
}

// Stop Stop the current fiber
func (fiber *Fiber) Stop() {
	if running == 1 {
		finished <- true
		running = 0
		stop <- true
		log.Plain("| Fiber Stopped |")
	} else {
		log.FiberWarning("No running fiber")
	}
}

// Run Run the current fiber
func (fiber *Fiber) Run() {
	running++
	if running > 1 {
		log.FiberWarning("A fiber is already running")
	} else {
		instruction := fiber.Instructions[0]
		variable.FiberVariable = nil

		for {
			select {
			case <-stop:
				return
			default:
				nextID := -1
				funcName := instruction.FuncName
				var instructionData reflect.Value
				if instruction.Datas != nil {
					instructionData = reflect.ValueOf(instruction.Datas).Elem()
				}
				switch instruction.Package {
				case "Flow":
					nextID = flowpack.Processing(funcName, instructionData, finished)
				// DON'T REMOVE ME / New processing inserted here
				case "Algorithmic":
					nextID = algorithmicpack.Processing(funcName, instructionData, finished)
				case "Api":
					nextID = apipack.Processing(funcName, instructionData, finished)
				case "Arithmetic":
					nextID = arithmeticpack.Processing(funcName, instructionData, finished)
				case "Array":
					nextID = arraypack.Processing(funcName, instructionData, finished)
				case "Battery":
					nextID = batterypack.Processing(funcName, instructionData, finished)
				case "Chronometer":
					nextID = chronometerpack.Processing(funcName, instructionData, finished)
				case "Clipboard":
					nextID = clipboardpack.Processing(funcName, instructionData, finished)
				case "Conversion":
					nextID = conversionpack.Processing(funcName, instructionData, finished)
				case "Define":
					nextID = definepack.Processing(funcName, instructionData, finished)
				case "Dictionary":
					nextID = dictionarypack.Processing(funcName, instructionData, finished)
				case "File":
					nextID = filepack.Processing(funcName, instructionData, finished)
				case "Input":
					nextID = inputpack.Processing(funcName, instructionData, finished)
				case "Json":
					nextID = jsonpack.Processing(funcName, instructionData, finished)
				case "Keyboard":
					nextID = keyboardpack.Processing(funcName, instructionData, finished)
				case "Log":
					nextID = logpack.Processing(funcName, instructionData, finished)
				case "Mouse":
					nextID = mousepack.Processing(funcName, instructionData, finished)
				case "Notification":
					nextID = notificationpack.Processing(funcName, instructionData, finished)
				case "Process":
					nextID = processpack.Processing(funcName, instructionData, finished)
				case "Scraping":
					nextID = scrapingpack.Processing(funcName, instructionData, finished)
				case "Screen":
					nextID = screenpack.Processing(funcName, instructionData, finished)
				case "Sleep":
					nextID = sleeppack.Processing(funcName, instructionData, finished)
				case "Sound":
					nextID = soundpack.Processing(funcName, instructionData, finished)
				case "Systime":
					nextID = systimepack.Processing(funcName, instructionData, finished)
				default:
					log.FiberError("Package not found:", instruction.Package)
					continue
				}

				// The error code for a finished fiber
				if nextID == finishedFiber {
					running = 0
					return
				}

				// The error code for an unknown function
				if nextID == unknownFunction {
					log.FiberError("Function not found:", funcName)
					nextID = continuFiber
				}

				if nextID == continuFiber {
					idx := sort.Search(len(fiber.Instructions), func(i int) bool {
						return fiber.Instructions[i].ID >= instruction.NextID
					})
					if idx == len(fiber.Instructions) {
						log.FiberFatalError("The instruction with the id", instruction.NextID, "has no been found")
						running = 0
						return
					} else {
						instruction = fiber.Instructions[idx]
					}
				} else {
					instruction = fiber.Instructions[nextID]
				}
			}
		}
	}
}

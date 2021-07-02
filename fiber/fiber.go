package fiber

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	// DON'T REMOVE ME / New packages inserted here

	"gotomate-astilectron/fiber/instructions"
	"gotomate-astilectron/fiber/packages"
	algorithmic "gotomate-astilectron/fiber/packages/Algorithmic"
	flow "gotomate-astilectron/fiber/packages/Flow"
	sleep "gotomate-astilectron/fiber/packages/Sleep"
	"gotomate-astilectron/fiber/template"
	"gotomate-astilectron/fiber/variable"
	"reflect"
	"sort"
)

// NewFiber Define the new automate's fiber
var NewFiber = new(Fiber)
var running = 0
var finished = make(chan bool)
var stop = make(chan bool)

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
	path := "saves/"
	if len(filePath) > 0 {
		path = filePath[0]
	}
	fullPath := path + fiber.Name + ".json"
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
		fmt.Println("GOTOMATE ERROR: Unable to open the saved fiber")
		return nil
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var loadingFiber LoadingFiber
	err = json.Unmarshal(byteValue, &loadingFiber)
	if err != nil {
		fmt.Println("GOTOMATE ERROR: Unable to extract the saved fiber")
	}

	fiber.Clean()
	fiber.Name = loadingFiber.Name

	for _, instruction := range loadingFiber.Instructions {

		newInstruction := &instructions.Instruction{
			ID:       instruction.ID,
			Package:  instruction.Package,
			FuncName: instruction.FuncName,
			X:        instruction.X,
			Y:        instruction.Y,
			IconPath: instruction.IconPath,
			NextID:   instruction.NextID,
		}
		databinder, fields := packages.PackageDecode(instruction.Package, instruction.FuncName)
		var temp = make([]template.Field, len(fields))
		if databinder != nil && fields != nil {

			if err := json.Unmarshal(instruction.Datas, databinder); err != nil {
				fmt.Println("GOTOMATE ERROR: Unable to convert the saved instruction")
			}

			if err := json.Unmarshal(instruction.Template, &temp); err != nil {
				fmt.Println("GOTOMATE ERROR: Unable to convert the saved template")
			}
			// newInstruction.Template = &temp
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
func (fiber *Fiber) Import(filePath string) {
	fiber.Open(filePath)
}

// Clean Delete all the instructions of the current fiber
func (fiber *Fiber) Clean() {
	p := reflect.ValueOf(fiber).Elem()
	p.Set(reflect.Zero(p.Type()))
}

// Stop Stop the current fiber
func (fiber *Fiber) Stop() {
	if running == 1 {
		fmt.Println("| Fiber Stopped |")
		finished <- true
		running = 0
		stop <- true
	} else {
		fmt.Println("FIBER WARNING: No running fiber")
	}
}

// Run Run the current fiber
func (fiber *Fiber) Run() {
	running++
	if running > 1 {
		fmt.Println("FIBER WARNING: A fiber is already running")
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
					ended := flow.Processing(funcName, instructionData, finished)
					if ended {
						running = 0
						return
					}
				// DON'T REMOVE ME / New processing inserted here
				case "Algorithmic":
					nextID = algorithmic.Processing(funcName, instructionData, finished)
				case "Sleep":
					nextID = sleep.Processing(funcName, instructionData, finished)
				default:
					fmt.Println("FIBER WARNING: This package is not integrated yet: " + instruction.Package)
					continue
				}

				if nextID == -1 {
					idx := sort.Search(len(fiber.Instructions), func(i int) bool {
						return fiber.Instructions[i].ID >= instruction.NextID
					})
					if idx == len(fiber.Instructions) {
						fmt.Println("FIBER FATAL ERROR: The instruction with the id", instruction.NextID, "has no been founded")
						fmt.Println("| Fiber Finished at Fatal Error |")
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
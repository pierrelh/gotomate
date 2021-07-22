package instructions

import (
	"encoding/json"
	"gotomate/fiber/packages"
	"gotomate/fiber/template"
	"gotomate/log"
	"reflect"
	"sort"

	"github.com/mitchellh/mapstructure"
)

// FiberInstructions Create the automate's list of instructions
type FiberInstructions []*Instruction

// Instruction Setting fiber's instructions structure
type Instruction struct {
	ID       int
	Package  string
	FuncName string
	X        int
	Y        int
	IconPath string
	NextID   int
	Datas    interface{}
	Template *template.InstructionTemplate
}

// FiberLoadingInstruction Create the automate's list of instructions
type FiberLoadingInstruction []*LoadingInstruction

// LoadingInstruction Initialize a loading fiber's instruction
type LoadingInstruction struct {
	ID       int
	Package  string
	FuncName string
	X        int
	Y        int
	IconPath string
	NextID   int
	Datas    json.RawMessage
}

// Hydrate the new Fiber's instructions
func (fi *FiberInstructions) Hydrate() *Instruction {
	databinder, fields := packages.PackageDecode("Flow", "Start")
	newInstruction := &Instruction{
		ID:       0,
		Package:  "Flow",
		FuncName: "Start",
		X:        0,
		Y:        0,
		IconPath: "../fiber/packages/Flow/icon.png",
		NextID:   1,
		Datas:    databinder,
		Template: fields,
	}
	*fi = append(*fi, newInstruction)
	return newInstruction
}

// Clean Delete all the automate's instructions
func (fi *FiberInstructions) Clean() {
	p := reflect.ValueOf(fi).Elem()
	p.Set(reflect.Zero(p.Type()))
}

// Delete delete a fiber instruction. Return true if the instruction was found, else return false
func (fi *FiberInstructions) DeleteInstruction(inst *Instruction) bool {
	for i := 0; i < len(*fi); i++ {
		if inst == (*fi)[i] {
			*fi = append((*fi)[:i], (*fi)[i+1:]...)
			return true
		}
	}
	log.GotomateError("Unable to delete the fiber's instruction")
	return false
}

// FindInstructionById get an instruction in the fiber's instructions by id
func (fi *FiberInstructions) FindInstructionById(id int) *Instruction {
	idx := sort.Search(len(*fi), func(i int) bool {
		return (*fi)[i].ID >= id
	})
	if idx == len(*fi) {
		log.GotomateError("Unable to find button with id:", id)
		return nil
	} else {
		return (*fi)[idx]
	}
}

// Build build a new instruction with the right parameters
func (inst *Instruction) Build(id int, content map[string]interface{}) *Instruction {
	databinder, fields := packages.PackageDecode(content["PackageName"].(string), content["FuncName"].(string))

	if databinder != nil && fields != nil {
		inst.Template = fields
		inst.Datas = databinder
	} else {
		inst.Template = nil
		inst.Datas = nil
	}
	funcName, packName := content["FuncName"].(string), content["PackageName"].(string)
	inst.ID = id
	inst.Package = packName
	inst.FuncName = funcName
	inst.IconPath = "../fiber/packages/" + packName + "/icon.png"
	inst.NextID = id + 1
	return inst
}

// UpdatePosition Update the position of an instruction in the datas
func (inst *Instruction) UpdatePosition(x, y int) {
	inst.X = x
	inst.Y = y
}

// UpdateNextID Update the NextID of an instruction in the datas
func (inst *Instruction) UpdateNextID(nextID int) {
	inst.NextID = nextID
}

func (inst *Instruction) UpdateDatabinder(databinder map[string]interface{}) {
	mapstructure.Decode(databinder, &inst.Datas)
}

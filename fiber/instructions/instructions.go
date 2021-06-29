package instructions

import (
	"encoding/json"
	"fmt"
	"gotomate-astilectron/fiber/packages"
	"gotomate-astilectron/fiber/template"
	"reflect"
	"sort"
	"strconv"
)

// https://anseki.github.io/leader-line/

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
	Template *template.Template
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
	Datas    interface{}
	Template json.RawMessage
}

// Hydrate the new Fiber's instructions
func (fi *FiberInstructions) Hydrate() *Instruction {
	databinder, fields := packages.PackageDecode("Flow", "Start")
	flds := make([]template.Field, len(fields))
	copy(flds, fields)
	temp := &template.Template{
		Fields: flds,
	}
	newInstruction := &Instruction{
		ID:       0,
		Package:  "Flow",
		FuncName: "Start",
		X:        0,
		Y:        0,
		IconPath: "../fiber/packages/Flow/icon.png",
		NextID:   1,
		Datas:    databinder,
		Template: temp,
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
	fmt.Println("GOTOMATE ERROR: Unable to delete the fiber's instruction")
	return false
}

// FindInstructionById get an instruction in the fiber's instructions by id
func (fi *FiberInstructions) FindInstructionById(id int) *Instruction {
	idx := sort.Search(len(*fi), func(i int) bool {
		return (*fi)[i].ID >= id
	})
	if idx == len(*fi) {
		fmt.Println("FIBER ERROR: Unable to find button with id: ", id)
		return nil
	} else {
		return (*fi)[idx]
	}
}

// Build build a new instruction with the right parameters
func (inst *Instruction) Build(id int, content map[string]interface{}) *Instruction {
	databinder, fields := packages.PackageDecode(content["PackageName"].(string), content["FuncName"].(string))

	if databinder != nil && fields != nil {
		flds := make([]template.Field, len(fields))
		copy(flds, fields)
		inst.Template = &template.Template{
			Fields: flds,
		}
	} else {
		inst.Template = nil
	}
	funcName, packName := content["FuncName"].(string), content["PackageName"].(string)
	inst.ID = id
	inst.Package = packName
	inst.FuncName = funcName
	inst.IconPath = "../fiber/packages/" + packName + "/icon.png"
	inst.NextID = id + 1
	inst.Datas = databinder
	return inst
}

// UpdatePosition update the position of an instruction in the datas
func (inst *Instruction) UpdatePosition(x, y int) {
	inst.X = x
	inst.Y = y
}

func getAttr(obj interface{}, fieldName string) reflect.Value {
	pointToStruct := reflect.ValueOf(obj)
	curStruct := pointToStruct.Elem()
	if curStruct.Kind() != reflect.Struct {
		panic("GOTOMATE ERROR: Struct not found in template")
	}
	curField := curStruct.FieldByName(fieldName)
	if !curField.IsValid() {
		panic("GOTOMATE ERROR: Field not found in template databinder")
	}
	return curField
}

func (inst *Instruction) SetDatabinder() {
	databinder := inst.Datas
	for i := 0; i < len(inst.Template.Fields); i++ {
		if inst.Template.Fields[i].Input != nil {
			isVar := false
			if inst.Template.Fields[i].VariableToggler != nil {
				varToggler := reflect.ValueOf(inst.Template.Fields[i].VariableToggler)
				isVar = varToggler.FieldByName("Checked").Interface().(bool)
				if bind := varToggler.FieldByName("Bind").String(); bind != "" {
					databinderInsert(getAttr(databinder, bind), isVar)
				}
			}
			input := reflect.ValueOf(inst.Template.Fields[i].Input)
			if bind := input.FieldByName("Bind").String(); bind != "" && !isVar {
				databinderInsert(getAttr(databinder, bind), input.FieldByName("Value").Interface())
			} else if bind := input.FieldByName("BindVariable").String(); bind != "" && isVar {
				databinderInsert(getAttr(databinder, bind), input.FieldByName("Value").Interface())
			}
		}
	}
	fmt.Printf("%+v\n", inst.Datas)
}

func databinderInsert(r reflect.Value, val interface{}) {
	switch r.Kind() {
	case reflect.Bool:
		r.SetBool(val.(bool))
	case reflect.Float32, reflect.Float64:
		f, err := strconv.ParseFloat(val.(string), 64)
		if err != nil {
			fmt.Println("GOTOMATE WARNING: Unable to parse data", val, "to float")
			return
		}
		r.SetFloat(f)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		i, err := strconv.ParseInt(val.(string), 10, 64)
		if err != nil {
			fmt.Println("GOTOMATE WARNING: Unable to parse data", val, "to int")
			return
		}
		r.SetInt(i)
	case reflect.String:
		r.SetString(val.(string))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		u, err := strconv.ParseUint(val.(string), 10, 64)
		if err != nil {
			fmt.Println("GOTOMATE WARNING: Unable to parse data", val, "to uint")
			return
		}
		r.SetUint(u)
	default:
		fmt.Println("GOTOMATE WARNING: Unknown type on setting datas")
	}
}

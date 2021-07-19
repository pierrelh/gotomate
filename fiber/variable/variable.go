package variable

import (
	"fmt"
	"gotomate-astilectron/log"
	"math"
	"reflect"
	"strconv"
)

// FiberVariable define the fiber's key / values
var FiberVariable []*InstructionVariable

// InstructionVariable Define a key / value of the fiber
type InstructionVariable struct {
	Key   string
	Value interface{}
}

// SearchVariable Search a key in fiber's keys / values array
func SearchVariable(name string) *InstructionVariable {
	idx := -1
	for i := 0; i < len(FiberVariable); i++ {
		if FiberVariable[i].Key == name {
			idx = i
			break
		}
	}
	if idx != -1 {
		return FiberVariable[idx]
	} else {
		log.FiberError("Unable to find the fiber's var: ", name)
		return &InstructionVariable{
			Value: nil,
		}
	}
}

// SetVariable create or update a key / value in fiber's variable
func SetVariable(key string, value interface{}) {
	for i := 0; i < len(FiberVariable); i++ {
		if FiberVariable[i].Key == key {
			FiberVariable[i].Value = value
			return
		}
	}
	newValue := &InstructionVariable{
		Key:   key,
		Value: value,
	}
	FiberVariable = append(FiberVariable, newValue)
}

// GetVariableType get the type of a value
func GetVariableType(value interface{}) string {
	switch value.(type) {
	case []bool:
		return "[]bool"
	case []complex64:
		return "[]complex64"
	case []complex128:
		return "[]complex128"
	case []float32:
		return "[]float32"
	case []float64:
		return "[]float64"
	case []int:
		return "[]int"
	case []int8:
		return "[]int8"
	case []int16:
		return "[]int16"
	case []int32:
		return "[]int32"
	case []int64:
		return "[]int64"
	case []string:
		return "[]string"
	case []uint:
		return "[]uint"
	case []uint8:
		return "[]uint8"
	case []uint16:
		return "[]uint16"
	case []uint32:
		return "[]uint32"
	case []uint64:
		return "[]uint64"
	case []uintptr:
		return "[]uintptr"
	case bool:
		return "bool"
	case complex64:
		return "complex64"
	case complex128:
		return "complex128"
	case float32:
		return "float32"
	case float64:
		return "float64"
	case int:
		return "int"
	case int8:
		return "int8"
	case int16:
		return "int16"
	case int32:
		return "int32"
	case int64:
		return "int64"
	case string:
		return "string"
	case uint:
		return "uint"
	case uint8:
		return "uint8"
	case uint16:
		return "uint16"
	case uint32:
		return "uint32"
	case uint64:
		return "uint64"
	case uintptr:
		return "uintptr"
	default:
		log.FiberWarning("Unknown variable type")
		return ""
	}
}

// GetValue find a value in fiber's values data
func GetValue(instructionData reflect.Value, varName string, args ...string) (interface{}, error) {
	if len(args) == 2 {
		if isAVar := instructionData.FieldByName(args[0]).Interface().(bool); isAVar {
			variableName := instructionData.FieldByName(varName).Interface().(string)
			if val := SearchVariable(variableName).Value; val != nil {
				return val, nil
			} else {
				log.FiberError("Variable not found")
				return false, fmt.Errorf("FIBER ERROR: Variable not found")
			}
		} else {
			return instructionData.FieldByName(args[1]).Interface(), nil
		}
	} else {
		variableName := instructionData.FieldByName(varName).Interface().(string)
		if val := SearchVariable(variableName).Value; val != nil {
			return val, nil
		} else {
			log.FiberError("Variable not found")
			return false, fmt.Errorf("FIBER ERROR: Variable not found")
		}
	}
}

// GetFloat Return the float value of an interface
func GetFloat(unk interface{}) float64 {
	var floatType = reflect.TypeOf(float64(0))
	var stringType = reflect.TypeOf("")
	switch i := unk.(type) {
	case float64:
		return i
	case float32:
		return float64(i)
	case int64:
		return float64(i)
	case int32:
		return float64(i)
	case int:
		return float64(i)
	case uint64:
		return float64(i)
	case uint32:
		return float64(i)
	case uint:
		return float64(i)
	case string:
		v, _ := strconv.ParseFloat(i, 64)
		return v
	default:
		v := reflect.ValueOf(unk)
		v = reflect.Indirect(v)
		if v.Type().ConvertibleTo(floatType) {
			fv := v.Convert(floatType)
			return fv.Float()
		} else if v.Type().ConvertibleTo(stringType) {
			sv := v.Convert(stringType)
			s := sv.String()
			v, _ := strconv.ParseFloat(s, 64)
			return v
		} else {
			return math.NaN()
		}
	}
}

// GetString Return the string value of an interface
func GetString(unk interface{}) string {
	str := fmt.Sprintf("%v", unk)
	return str
}

// ToArrayOfString Return the array of string value of an array of interface
func ToArrayOfString(unk []interface{}) []string {
	var strings []string
	for _, element := range unk {
		strings = append(strings, GetString(element))
	}
	return strings
}

package variable

import (
	"gotomate/log"
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
func SearchVariable(name string) (interface{}, error) {
	idx := -1
	for i := 0; i < len(FiberVariable); i++ {
		if FiberVariable[i].Key == name {
			idx = i
			break
		}
	}
	if idx != -1 {
		return FiberVariable[idx].Value, nil
	} else {
		log.FiberError("Unable to find the fiber's var: ", name)
		return nil, log.Error("FIBER ERROR: Variable not found")
	}
}

// SetVariable create or update a key / value in fiber's variable
func SetVariable(object interface{}, objectkey string, value interface{}) {
	key := object.(reflect.Value).FieldByName(objectkey).Interface().(string)

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

// Keys set the databinder's fields to search trough
type Keys struct {
	Name      string // The name of the raw value
	IsVarName string // The name of the value that bind if the value is a var
	VarName   string // The name of the var value
}

// GetValue find a value in fiber's values data
func (k Keys) GetValue(data interface{}) (interface{}, error) {
	datas := data.(reflect.Value)

	if k.Name == "" && k.VarName == "" {
		return nil, log.Error("FIBER ERROR: The set of data is incorrect in GetValue")
	} else {
		if k.Name != "" && k.IsVarName == "" && k.VarName == "" {
			return datas.FieldByName(k.Name).Interface(), nil
		} else if k.VarName != "" && k.IsVarName == "" && k.Name == "" {
			return SearchVariable(datas.FieldByName(k.VarName).Interface().(string))
		} else {
			if isAVar := datas.FieldByName(k.IsVarName).Interface().(bool); isAVar {
				return SearchVariable(datas.FieldByName(k.VarName).Interface().(string))
			} else {
				return datas.FieldByName(k.Name).Interface(), nil
			}
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
	str := log.Sprint("%v", unk)
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

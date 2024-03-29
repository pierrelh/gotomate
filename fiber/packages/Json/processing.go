package json

// Processing process the functions from clipboard's package
func Processing(funcName string, instructionData interface{}, finished chan bool) int {
	nextID := -1
	switch funcName {
	case "CreateJson":
		go func() {
			nextID = CreateJson(instructionData, finished)
		}()
		<-finished
	case "JsonToDictionary":
		go func() {
			nextID = JsonToDictionary(instructionData, finished)
		}()
		<-finished
	case "SaveJson":
		go func() {
			nextID = SaveJson(instructionData, finished)
		}()
		<-finished
	default:
		return -2
	}
	return nextID
}

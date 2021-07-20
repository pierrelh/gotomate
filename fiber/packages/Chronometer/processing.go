package chronometer

// Processing process the functions from clipboard's package
func Processing(funcName string, instructionData interface{}, finished chan bool) int {
	nextID := -1
	switch funcName {
	case "End":
		go func() {
			nextID = End(instructionData, finished)
		}()
		<-finished
	case "Start":
		go func() {
			nextID = Start(instructionData, finished)
		}()
		<-finished
	default:
		return -2
	}
	return nextID
}

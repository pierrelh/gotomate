package log

// Processing process the functions from log's package
func Processing(funcName string, instructionData interface{}, finished chan bool) int {
	nextID := -1
	switch funcName {
	case "Print":
		go func() {
			nextID = Print(instructionData, finished)
		}()
		<-finished
	default:
		return -2
	}
	return nextID
}

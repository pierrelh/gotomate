package notification

// Processing process the functions from notification's package
func Processing(funcName string, instructionData interface{}, finished chan bool) int {
	nextID := -1
	switch funcName {
	case "Create":
		go func() {
			nextID = Create(instructionData, finished)
		}()
		<-finished
	default:
		return -2
	}
	return nextID
}

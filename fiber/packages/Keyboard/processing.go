package keyboard

// Processing process the functions from keyboard's package
func Processing(funcName string, instructionData interface{}, finished chan bool) int {
	nextID := -1
	switch funcName {
	case "Tap":
		go func() {
			nextID = Tap(instructionData, finished)
		}()
		<-finished
	case "Type":
		go func() {
			nextID = Type(instructionData, finished)
		}()
		<-finished
	default:
		return -2
	}
	return nextID
}

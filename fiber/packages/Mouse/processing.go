package mouse

// Processing process the functions from mouse's package
func Processing(funcName string, instructionData interface{}, finished chan bool) int {
	nextID := -1
	switch funcName {
	case "Click":
		go func() {
			nextID = Click(instructionData, finished)
		}()
		<-finished
	case "Drag":
		go func() {
			nextID = Drag(instructionData, finished)
		}()
		<-finished
	case "DragSmooth":
		go func() {
			nextID = DragSmooth(instructionData, finished)
		}()
		<-finished
	case "Move":
		go func() {
			nextID = Move(instructionData, finished)
		}()
		<-finished
	case "MoveSmooth":
		go func() {
			nextID = MoveSmooth(instructionData, finished)
		}()
		<-finished
	case "Scroll":
		go func() {
			nextID = Scroll(instructionData, finished)
		}()
		<-finished
	default:
		return -2
	}
	return nextID
}

package api

// Processing process the functions from clipboard's package
func Processing(funcName string, instructionData interface{}, finished chan bool) int {
	nextID := -1
	switch funcName {
	case "Get":
		go func() {
			nextID = Get(instructionData, finished)
		}()
		<-finished
	case "Post":
		go func() {
			nextID = Post(instructionData, finished)
		}()
		<-finished
	default:
		return -2
	}
	return nextID
}

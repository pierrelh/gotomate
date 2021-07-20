package arithmetic

// Processing process the functions from clipboard's package
func Processing(funcName string, instructionData interface{}, finished chan bool) int {
	nextID := -1
	switch funcName {
	case "Divide":
		go func() {
			nextID = Divide(instructionData, finished)
		}()
		<-finished
	case "Multiply":
		go func() {
			nextID = Multiply(instructionData, finished)
		}()
		<-finished
	case "Pow":
		go func() {
			nextID = Pow(instructionData, finished)
		}()
		<-finished
	case "Sqrt":
		go func() {
			nextID = Sqrt(instructionData, finished)
		}()
		<-finished
	case "Substract":
		go func() {
			nextID = Substract(instructionData, finished)
		}()
		<-finished
	case "Sum":
		go func() {
			nextID = Sum(instructionData, finished)
		}()
		<-finished
	default:
		return -2
	}
	return nextID
}

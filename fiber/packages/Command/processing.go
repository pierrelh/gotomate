package command

// Processing process the functions from clipboard's package
func Processing(funcName string, instructionData interface{}, finished chan bool) int {
	nextID := -1
	switch funcName {
	case "Dir":
		go func() {
			nextID = Dir(instructionData, finished)
		}()
		<-finished
	case "GetOutput":
		go func() {
			nextID = GetOutput(instructionData, finished)
		}()
		<-finished
	case "New":
		go func() {
			nextID = New(instructionData, finished)
		}()
		<-finished
	case "Run":
		go func() {
			nextID = Run(instructionData, finished)
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

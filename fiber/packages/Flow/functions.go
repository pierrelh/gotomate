package flow

import "fmt"

// Write write string to clipboard
func End(finished chan bool) {
	fmt.Println("| Fiber Finished |")
	finished <- true
}

// Read read string from clipboard
func Start(finished chan bool) {
	fmt.Println("| Fiber Start |")
	finished <- true
}

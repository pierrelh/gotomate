package flow

import (
	"gotomate/log"
)

// Write write string to clipboard
func End(finished chan bool) {
	log.Plain("| Fiber Finished |")
	finished <- true
}

// Read read string from clipboard
func Start(finished chan bool) {
	log.Plain("| Fiber Start |")
	finished <- true
}

package log

import (
	"fmt"
	"os"
)

func FiberScan(target ...interface{}) {
	fmt.Println("SCANING:")
	fmt.Scanln(target)
}

// FiberInfo log an info on the fiber's running
func FiberInfo(msg ...interface{}) {
	fmt.Println("FIBER INFO:", msg)
}

// FiberError log an error on the fiber's running
func FiberError(msg ...interface{}) {
	fmt.Println("FIBER ERROR:", msg)
}

// FiberWarning log a warning on the fiber's running
func FiberWarning(msg ...interface{}) {
	fmt.Println("FIBER WARNING:", msg)
}

// FiberFatalError log a fatal error on the fiber's running
func FiberFatalError(msg ...interface{}) {
	fmt.Println("FIBER FATAL ERROR:", msg)
	fmt.Println("| Fiber Finished at Fatal Error |")
}

// GotomateError log an error on Gotomate process
func GotomateError(msg ...interface{}) {
	fmt.Println("GOTOMATE ERROR:", msg)
}

// GotomateFatalError log a fatal error on Gotomate process & exit the process
func GotomateFatalError(msg ...interface{}) {
	fmt.Println("GOTOMATE FATAL ERROR:", msg)
	os.Exit(1)
}

// Plain log a plain message
func Plain(msg ...interface{}) {
	fmt.Println(msg...)
}

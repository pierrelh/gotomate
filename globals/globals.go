package globals

import "os"

const (
	AppName string = "Gotomate"
	AppIcon string = "/resources/gotomate.png"
)

var (
	DirectoryPath, _ = os.Getwd()
)

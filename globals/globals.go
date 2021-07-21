package globals

import "os"

const (
	AppName string = "Gotomate"
	AppIcon string = "/img/gotomate.png"
)

var (
	DirectoryPath, _ = os.Getwd()
)

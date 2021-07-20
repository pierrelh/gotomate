package globals

import "os"

const (
	AppName string = "Gotomate"
)

var (
	DirectoryPath, _ = os.Getwd()
	AppIcon          = DirectoryPath + "/img/gotomate.png"
	AppTemplate      = DirectoryPath + "/app/template.html"
)

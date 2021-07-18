package files

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type File struct {
	IsDir     bool
	Name      string
	Extension string
}

// GetHomeJson Get the required filetypes & the directories of the selected path directory
func GetHomeJson(path, filetype string) ([]*File, string) {
	var content []*File
	var err error

	// Select the home folder if the path is not explicit
	if path == "home" {
		path, err = os.UserHomeDir()

		if err != nil {
			fmt.Println("GOTOMATE ERROR: Unable to find Home directory")
			return nil, ""
		}
	}

	files, err := ioutil.ReadDir(path)

	if err != nil {
		fmt.Println("GOTOMATE ERROR: Unable to read directory", path)
		return nil, ""
	}
	for _, f := range files {
		if f.IsDir() || (filetype != "none" && filepath.Ext(f.Name()) == filetype) {
			file := &File{
				IsDir:     f.IsDir(),
				Name:      f.Name(),
				Extension: filepath.Ext(f.Name()),
			}
			content = append(content, file)
		}
	}
	fmt.Println(content)
	return content, path
}

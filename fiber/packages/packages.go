package packages

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

type GotomatePackages struct {
	Packages []*Package
}

// Package Declare the structure of a package
type Package struct {
	Name      string
	Functions []string
}

// Hydrate Fill the GotomatePackages structure with current availables packages
func (gp *GotomatePackages) Hydrate() *GotomatePackages {
	const subPackage = "fiber/packages"

	filepath.Walk(subPackage, func(path string, info os.FileInfo, err error) error {
		if filepath.Base(path) == "functions.go" {
			pathSlice := strings.Split(filepath.Dir(path), "\\")
			directory := pathSlice[len(pathSlice)-1]
			newPackage := &Package{
				Name: directory,
			}
			set := token.NewFileSet()
			pack, err := parser.ParseFile(set, path, nil, 0)
			if err != nil {
				fmt.Println("GOTOMATE FATAL ERROR: Unable to parse the file", err)
				os.Exit(1)
			}

			for _, d := range pack.Decls {
				if fn, isFn := d.(*ast.FuncDecl); isFn {
					newPackage.Functions = append(newPackage.Functions, fn.Name.String())
				}
			}
			gp.Packages = append(gp.Packages, newPackage)
		}
		return nil
	})
	return gp
}

func InitImportPackage() {

}

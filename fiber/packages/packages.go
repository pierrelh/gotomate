package packages

import (
	"archive/zip"
	"go/ast"
	"go/parser"
	"go/token"
	"gotomate-astilectron/log"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
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
				log.GotomateFatalError("Unable to parse the file", err)
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

// ImportPackage Import a new package to Gotomate
func ImportPackage(path string) error {

	r, err := zip.OpenReader(path)
	if err != nil {
		log.GotomateError("Unable to open the zip reader")
		return err
	}
	defer r.Close()

	validation, dirs := ScanArchive(r.File)

	if validation {

		for _, f := range r.File {

			// Store filename/path for returning and using later on
			fpath := filepath.Join("./fiber/packages", f.Name)

			// Check for ZipSlip. More Info: http://bit.ly/2MsjAWE
			if !strings.HasPrefix(fpath, filepath.Clean("./fiber/packages")+string(os.PathSeparator)) {
				log.GotomateError("Illegal file path given")
				return log.Error("%s: illegal file path", fpath)
			}

			if f.FileInfo().IsDir() {
				// Make Folder
				os.MkdirAll(fpath, os.ModePerm)
				continue
			}

			// Make File
			if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
				log.GotomateError("Unable to create the new directory")
				return err
			}

			outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				log.GotomateError("Unable to open the new file")
				return err
			}

			rc, err := f.Open()
			if err != nil {
				log.GotomateError("Unable to open the read closer")
				return err
			}

			_, err = io.Copy(outFile, rc)

			// Close the file without defer to close before next iteration of loop
			outFile.Close()
			rc.Close()

			if err != nil {
				log.GotomateError("Unable to copy the new files")
				return err
			}
		}
	}

	for i := 0; i < len(dirs); i++ {
		dirs[i] = strings.TrimSuffix(dirs[i], "/")
		packageName := strings.ToLower(dirs[i])
		packageImport := "// DON'T REMOVE ME / New packages inserted here\n" + packageName + " \"gotomate/fiber/packages/" + dirs[i] + "\""
		processing := "// DON'T REMOVE ME / New processing inserted here\n" + "case \"" + dirs[i] + "\":\n nextID = " + packageName + ".Processing(funcName, instructionData, finished)"

		// Write in fiber.go (import the package & insert the Processing function)
		content, err := ioutil.ReadFile("./fiber/fiber.go")
		if err != nil {
			log.GotomateError("Unable to read the file fiber.go")
			return err
		}
		SContent := string(content)
		SContent = strings.Replace(SContent, "// DON'T REMOVE ME / New packages inserted here", packageImport, 1)
		SContent = strings.Replace(SContent, "// DON'T REMOVE ME / New processing inserted here", processing, 1)
		content = []byte(SContent)

		err = ioutil.WriteFile("./fiber/fiber.go", content, 0644)
		if err != nil {
			log.GotomateError("Unable to write the file fiber.go")
			return err
		}

		// Write in packages-dialog.go (import the package & insert the Build function)
		build := "// DON'T REMOVE ME / New Build inserted here\n" + "case \"" + dirs[i] + "\":\n databinder, template = " + packageName + ".Build(funcName)"

		content, err = ioutil.ReadFile("./fiber/packages/packages-dialog.go")
		if err != nil {
			log.GotomateError("Unable to read the file packages-dialog.go")
			return err
		}
		SContent = string(content)
		SContent = strings.Replace(SContent, "// DON'T REMOVE ME / New packages inserted here", packageImport, 1)
		SContent = strings.Replace(SContent, "// DON'T REMOVE ME / New Build inserted here", build, 1)
		content = []byte(SContent)

		err = ioutil.WriteFile("./fiber/packages/packages-dialog.go", content, 0644)
		if err != nil {
			log.GotomateError("Unable to write the file packages-dialog.go")
			return err
		}
	}
	return nil
}

// ScanArchive the archive to check if the format is correct
func ScanArchive(archive []*zip.File) (bool, []string) {
	var dir []string
	var filenames = [6]string{"build.go", "databinders.go", "functions.go", "icon.png", "processing.go", "templates.go"}

	for _, f := range archive {
		if f.FileInfo().IsDir() {
			dir = append(dir, f.Name)
		} else if len(dir) != 0 {
			for i := 0; i < len(dir); i++ {
				if matched, _ := regexp.Match(dir[i]+".*", []byte(f.Name)); matched {
					for y := 0; y < len(filenames); y++ {
						if matched, _ := regexp.Match(filenames[i]+".*", []byte(f.Name)); matched {
							break
						} else if i == len(filenames)-1 {
							log.GotomateError("Unknow file found in the archive", f.Name)
							return false, dir
						}
					}
				} else if i == len(dir)-1 {
					log.GotomateError("No files allowed at the archive root", f.Name)
					return false, dir
				}

			}
		} else {
			log.GotomateError("No files allowed at the archive root", f.Name)
			return false, dir
		}
	}
	return true, dir
}

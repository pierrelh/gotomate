package app

import (
	"fmt"
	"gotomate-astilectron/app/message"
	"gotomate-astilectron/fiber"
	"gotomate-astilectron/fiber/packages"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/asticode/go-astikit"
	"github.com/asticode/go-astilectron"
	"github.com/gonutz/w32"
)

var err error

// App define the structure of an app
type App struct {
	Asti    *astilectron.Astilectron
	Log     *log.Logger
	Options astilectron.Options
	Window  *astilectron.Window
	Menu    *astilectron.Menu
}

// Build build a new app
func (a *App) Build() *App {
	a = &App{
		Log: log.New(log.Writer(), log.Prefix(), log.Flags()),
		Options: astilectron.Options{
			AppName:            "Gotomate",
			BaseDirectoryPath:  "gotomate-astilectron",
			AppIconDefaultPath: "gotomate.png",
		},
		Window: new(astilectron.Window),
	}

	a.New()
	a.Start()

	a.Window, err = a.Asti.NewWindow(
		"./app/template.html",
		&astilectron.WindowOptions{
			Center: astikit.BoolPtr(true),
			Height: astikit.IntPtr(700),
			Width:  astikit.IntPtr(700),
		},
	)
	if err != nil {
		a.Log.Fatal(fmt.Errorf("main: new window failed: %w", err))
	}
	// Handle signals
	return a
}

// New init a new astilectron instance
func (a *App) New() {
	if a.Asti, err = astilectron.New(a.Log, a.Options); err != nil {
		a.Log.Fatal(fmt.Errorf("main: creating astilectron failed: %w", err))
	}
}

// Start start astilectron
func (a *App) Start() {
	a.Asti.HandleSignals()
	if err := a.Asti.Start(); err != nil {
		a.Log.Fatal(fmt.Errorf("main: starting astilectron failed: %w", err))
	}
}

// Create create the astilectron window
func (a *App) Create() {
	if err := a.Window.Create(); err != nil {
		a.Log.Fatal(fmt.Errorf("main: creating window failed: %w", err))
	}
}

// UpdateSavedFibersMenu Add all the saved fibers to the My fibers's menu
func (a *App) UpdateSavedFibersMenu() []*astilectron.MenuItemOptions {
	root := "./saves"
	var subMenus []*astilectron.MenuItemOptions

	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".json" {
			fullPath, _ := filepath.Abs(path)
			var file = filepath.Base(path)
			var extension = filepath.Ext(file)
			var name = file[0 : len(file)-len(extension)]
			subMenu := &astilectron.MenuItemOptions{
				Label: astikit.StrPtr(name),
				OnClick: func(e astilectron.Event) (deleteListener bool) {
					fmt.Println(fullPath)
					return
				},
			}
			subMenus = append(subMenus, subMenu)
		}
		return nil
	})

	// var ni = a.Menu.NewItem(&astilectron.MenuItemOptions{
	// 	Label:   astikit.StrPtr("My Fibers"),
	// 	SubMenu: subMenus,
	// })
	// s, _ := a.Menu.SubMenu(0)
	// ms, _ := s.Item(3)
	// ms.SubMenu() = ni
	// // delete ms.SubMenu() = ni
	// s.Append(3, ni)
	return subMenus
}

// CreateMenu create the menu for the astilectron window
func (a *App) CreateMenu() {
	if a.Menu != nil {
		a.Menu.Destroy()
	}
	a.Window.NewMenu([]*astilectron.MenuItemOptions{
		{
			Label: astikit.StrPtr("File"),
			SubMenu: []*astilectron.MenuItemOptions{
				{
					Label: astikit.StrPtr("New"),
					OnClick: func(e astilectron.Event) (deleteListener bool) {
						if fiber.NewFiber.IsSaved() {
							a.Window.SendMessage(
								message.New("InitFiber", fiber.NewFiber.New()),
							)
						} else {
							a.Window.SendMessage(
								message.New("InitNewFiber", nil),
							)
						}
						return
					}},
				{
					Label: astikit.StrPtr("Import"),
					OnClick: func(e astilectron.Event) (deleteListener bool) {
						fiber.NewFiber.InitImport()
						return
					},
				},
				{
					Label: astikit.StrPtr("Export"),
					OnClick: func(e astilectron.Event) (deleteListener bool) {
						fiber.NewFiber.Export()
						return
					},
				},
				{
					Label:   astikit.StrPtr("My Fibers"),
					SubMenu: a.UpdateSavedFibersMenu(),
				},
				{
					Label: astikit.StrPtr("Save"),
					OnClick: func(e astilectron.Event) (deleteListener bool) {
						if fiber.NewFiber.Name != "" {
							if !fiber.NewFiber.IsSaved() {
								fiber.NewFiber.Save()
								a.CreateMenu()
							}
						} else {
							a.Window.SendMessage(
								message.New("FiberUnamed", nil),
							)
						}
						return
					},
				},
				{
					Label: astikit.StrPtr("Exit"),
					OnClick: func(e astilectron.Event) (deleteListener bool) {
						a.Asti.Close()
						return
					},
				},
			},
		},
		{
			Label: astikit.StrPtr("Import Package"),
			OnClick: func(e astilectron.Event) (deleteListener bool) {
				packages.InitImportPackage()
				return
			},
		},
		{
			Label: astikit.StrPtr("Fiber"),
			SubMenu: []*astilectron.MenuItemOptions{
				{
					Label: astikit.StrPtr("Run"),
					OnClick: func(e astilectron.Event) (deleteListener bool) {
						go fiber.NewFiber.Run()
						return
					},
				},
				{
					Label: astikit.StrPtr("Stop"),
					OnClick: func(e astilectron.Event) (deleteListener bool) {
						go fiber.NewFiber.Stop()
						return
					},
				},
				{
					Label: astikit.StrPtr("Show Console"),
					Type:  astilectron.MenuItemTypeCheckbox,
					OnClick: func(e astilectron.Event) (deleteListener bool) {
						console := w32.GetConsoleWindow()
						if console != 0 {
							_, consoleProcID := w32.GetWindowThreadProcessId(console)
							if *e.MenuItemOptions.Checked {
								if w32.GetCurrentProcessId() == consoleProcID {
									w32.ShowWindowAsync(console, w32.SW_HIDE)
								}
							} else {
								if w32.GetCurrentProcessId() == consoleProcID {
									w32.ShowWindowAsync(console, w32.SW_SHOW)
								}
							}
						}
						return
					},
				},
			},
		},
		{
			Label: astikit.StrPtr("Help"),
			SubMenu: []*astilectron.MenuItemOptions{
				{
					Label: astikit.StrPtr("Documentation"),
					OnClick: func(e astilectron.Event) (deleteListener bool) {
						exec.Command("cmd", "/c", "start", "https://production-gotomate.herokuapp.com/documentation").Start()
						return
					},
				},
				{
					Label: astikit.StrPtr("Packages"),
					OnClick: func(e astilectron.Event) (deleteListener bool) {
						exec.Command("cmd", "/c", "start", "https://production-gotomate.herokuapp.com/packages").Start()
						return
					},
				},
			},
		},
	}).Create()
}

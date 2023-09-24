package main

import (
	"embed"
	"fmt"
	"os/exec"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"

	"github.com/jianxinliu/mydata/misc"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  fmt.Sprintf("mydata(%s)", misc.VERSION),
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		OnShutdown:       app.shutdown,
		Bind: []interface{}{
			app,
		},
		Menu:               createMenu(),
		Logger:             misc.AppLogger,
		LogLevel:           logger.DEBUG,
		LogLevelProduction: logger.INFO,
		Debug: options.Debug{
			OpenInspectorOnStartup: misc.IsDev(),
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

func createMenu() *menu.Menu {
	appMenu := menu.NewMenu()
	appMenu.Append(menu.AppMenu())
	appMenu.Append(menu.EditMenu())
	appMenu.Append(menu.WindowMenu())

	systemMenu := appMenu.AddSubmenu("系统")
	systemMenu.AddText("查看日志", keys.CmdOrCtrl("L"), func(_ *menu.CallbackData) {
		cmd := exec.Command("open", misc.LOG_PATH)
		cmd.Run()
	})
	if misc.IsDev() {
		systemMenu.AddText("查看store", keys.CmdOrCtrl("F"), func(_ *menu.CallbackData) {
			cmd := exec.Command("open", misc.STORE_FILE)
			cmd.Run()
		})
	}
	return appMenu
}

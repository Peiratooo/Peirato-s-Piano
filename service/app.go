package service

import (
	"embed"
	"github.com/wailsapp/wails/v3/pkg/application"
	"log"
	"os/exec"
	"syscall"
)

var (
	App        *application.App
	PianoWin   *application.WebviewWindow
	PRODUCTION = false
)

type WindowSize struct {
	Width  int
	Height int
}

func initSize() WindowSize {
	var res WindowSize
	w, _, _ := syscall.NewLazyDLL(`User32.dll`).NewProc(`GetSystemMetrics`).Call(uintptr(0))
	res.Width = int(w) * 17 / 20
	if res.Width < 900 {
		res.Width = 1800
	}
	res.Height = res.Width * 2 / 15
	return res
}

func Run(assets embed.FS) {
	size := initSize()
	App = application.New(application.Options{
		Name:        "Peirato's Piano",
		Description: "A piano keyboard on Windows",
		Services: []application.Service{
			application.NewService(&Keyboard{}),
		},
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(assets),
		},
		SingleInstance: &application.SingleInstanceOptions{
			UniqueID:               "com.peirato.piano",
			OnSecondInstanceLaunch: nil,
			AdditionalData:         nil,
			ExitCode:               0,
		},
		LogLevel:   application.DialogWarning,
		OnShutdown: CloseMidiDevice,
	})
	PianoWin = App.Window.NewWithOptions(application.WebviewWindowOptions{
		Title: "Peirato's Piano",
		Mac: application.MacWindow{
			InvisibleTitleBarHeight: 50,
			TitleBar:                application.MacTitleBarHidden,
		},
		Windows: application.WindowsWindow{
			BackdropType:                      0,
			WindowMaskDraggable:               true,
			DisableFramelessWindowDecorations: true,
		},

		Frameless:      true,
		Width:          size.Width,
		Height:         size.Height,
		MinHeight:      100,
		MinWidth:       800,
		BackgroundType: application.BackgroundTypeTranslucent,

		BackgroundColour:       application.NewRGBA(0, 0, 0, 0),
		URL:                    "/",
		DevToolsEnabled:        !PRODUCTION,
		OpenInspectorOnStartup: !PRODUCTION,
		EnableDragAndDrop:      true,
		AlwaysOnTop:            true,
	})
	App.Window.Add(PianoWin)

	err := App.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func (k *Keyboard) OpenUrl(url string) {
	exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Run()
}

func (k *Keyboard) Quit() {
	App.Quit()
}

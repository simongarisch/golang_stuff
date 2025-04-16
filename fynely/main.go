// fyne serve
// curl https://wasmtime.dev/install.sh -sSf | bash
// GOOS=js GOARCH=wasm go build -o main.wasm main.go
// fyne package -os web
package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func makeUI() (*widget.Label, *widget.Entry) {
	return widget.NewLabel("Hello world!"), widget.NewEntry()
}

func updateTime(clock *widget.Label) {
	formatted := time.Now().Format("Time: 03:04:05")
	clock.SetText(formatted)
}

func main() {
	a := app.New()

	// the clock demo
	w := a.NewWindow("Clock")

	clock := widget.NewLabel("")
	updateTime(clock)

	go func() {
		for range time.Tick(time.Second) {
			updateTime(clock)
		}
	}()

	label, entry := makeUI()
	vbox := container.NewVBox(clock, label, entry)
	w.SetContent(vbox)
	w.Show()

	// and another window
	w2 := a.NewWindow("Larger")
	w2.SetContent(widget.NewLabel("More content"))
	w2.Resize(fyne.NewSize(100, 100))
	w2.Show()

	a.Run()
}
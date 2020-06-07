package main

import (
	"log"
	"strconv"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	para := []*widgets.Paragraph{}
	for i := 1; i <= 4; i++ {
		p := widgets.NewParagraph()
		p.Text = "Section " + strconv.Itoa(i)
		para = append(para, p)
	}

	grid := ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, termHeight)

	grid.Set(
		ui.NewRow(1.0/2,
			ui.NewCol(0.5, para[0]),
			ui.NewCol(0.5, para[1]),
		),
		ui.NewRow(1.0/2,
			ui.NewCol(0.5, para[2]),
			ui.NewCol(0.5, para[3]),
		),
	)

	ui.Render(grid)
	uiEvents := ui.PollEvents()

	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				return
			case "<Resize>":
				payload := e.Payload.(ui.Resize)
				grid.SetRect(0, 0, payload.Width, payload.Height)
				ui.Clear()
				ui.Render(grid)
			}
		}
	}
}

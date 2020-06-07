package main

import (
	"log"
	"math"
	"strconv"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	// get the paragraphs
	para := []*widgets.Paragraph{}
	for i := 1; i <= 2; i++ {
		p := widgets.NewParagraph()
		p.Text = "Section " + strconv.Itoa(i)
		para = append(para, p)
	}

	// a pie chart
	pc := widgets.NewPieChart()
	pc.Title = "Pie Chart"
	pc.SetRect(5, 5, 70, 36)
	pc.Data = []float64{.25, .25, .25, .25}
	pc.AngleOffset = -.5 * math.Pi

	// and a bar graph
	sbc := widgets.NewStackedBarChart()
	sbc.Title = "Student's Marks: X-Axis=Name, Y-Axis=Grade% (Math, English, Science, Computer Science)"
	sbc.Labels = []string{"Ken", "Rob", "Dennis", "Linus"}

	sbc.Data = make([][]float64, 4)
	sbc.Data[0] = []float64{90, 85, 90, 80}
	sbc.Data[1] = []float64{70, 85, 75, 60}
	sbc.Data[2] = []float64{75, 60, 80, 85}
	sbc.Data[3] = []float64{100, 100, 100, 100}
	sbc.BarWidth = 5
	sbc.BorderStyle.Fg = ui.ColorCyan

	// place this in a grid
	grid := ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, termHeight)

	grid.Set(
		ui.NewRow(1.0/2,
			ui.NewCol(0.5, para[0]),
			ui.NewCol(0.5, para[1]),
		),
		ui.NewRow(1.0/2,
			ui.NewCol(0.5, pc),
			ui.NewCol(0.5, sbc),
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

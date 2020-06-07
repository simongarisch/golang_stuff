package main

import (
	"log"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

// TableStruct - We can select the table rows
type TableStruct struct {
	Table     *widgets.Table
	ActiveRow int
}

// NewTableStruct - Set the basic data
func NewTableStruct() TableStruct {
	table := widgets.NewTable()
	table.Rows = [][]string{
		[]string{"header1", "header2", "header3"},
		[]string{"AAA", "BBB", "CCC"},
		[]string{"DDD", "EEE", "FFF"},
		[]string{"GGG", "HHH", "III"},
	}
	table.TextStyle = ui.NewStyle(ui.ColorWhite)
	table.RowSeparator = true
	table.BorderStyle = ui.NewStyle(ui.ColorGreen)
	table.FillRow = true

	tableStruct := TableStruct{
		Table:     table,
		ActiveRow: 1,
	}
	tableStruct.setActiveRow()
	return tableStruct
}

func (tableStruct *TableStruct) setActiveRow() {
	table := tableStruct.Table
	numRows := len(table.Rows)
	table.RowStyles[0] = ui.NewStyle(ui.ColorWhite, ui.ColorBlack, ui.ModifierBold)
	i := tableStruct.ActiveRow

	for row := 1; row < numRows; row++ {
		if row == i {
			table.RowStyles[row] = ui.NewStyle(ui.ColorWhite, ui.ColorRed, ui.ModifierBold)
		} else {
			table.RowStyles[row] = ui.NewStyle(ui.ColorWhite, ui.ColorBlack, ui.ModifierBold)
		}
	}
}

// DownPressed - User selects the next row down
func (tableStruct *TableStruct) DownPressed() {
	numRows := len(tableStruct.Table.Rows)
	tableStruct.ActiveRow++
	if tableStruct.ActiveRow >= numRows {
		tableStruct.ActiveRow = numRows - 1
	}
	tableStruct.setActiveRow()
}

// UpPressed - User selects the next row up
func (tableStruct *TableStruct) UpPressed() {
	tableStruct.ActiveRow--
	if tableStruct.ActiveRow <= 0 {
		tableStruct.ActiveRow++
	}
	tableStruct.setActiveRow()
}

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	// create a table
	tableStruct := NewTableStruct()

	// create a grid
	grid := ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, termHeight)

	grid.Set(
		ui.NewRow(1.0,
			ui.NewCol(1.0, tableStruct.Table),
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
			case "w":
				tableStruct.UpPressed()
				ui.Clear()
				ui.Render(grid)
			case "s":
				tableStruct.DownPressed()
				ui.Clear()
				ui.Render(grid)
			case "<Resize>":
				payload := e.Payload.(ui.Resize)
				grid.SetRect(0, 0, payload.Width, payload.Height)
				ui.Clear()
				ui.Render(grid)
			}
		}
	}
}

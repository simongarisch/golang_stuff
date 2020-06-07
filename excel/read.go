package main

import "github.com/360EntSecGroup-Skylar/excelize"

func main() {
	f, err := excelize.OpenFile("GoBook.xlsx")
	if err != nil {
		println(err.Error())
		return
	}

	// Get value from cell by given worksheet name and axis.
	cell := f.GetCellValue("Blah", "B2")
	println(cell)

	// Get all the rows in the Sheet1.
	rows := f.GetRows("Blah")
	for _, row := range rows {
		for _, colCell := range row {
			print(colCell, "\t")
		}
		println()
	}
}

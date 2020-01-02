package main

import (
	"os"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	f := excelize.NewFile()

	// Create a new sheet
	index := f.NewSheet("Blah")

	// Set value of a cell
	f.SetCellValue("Blah", "A2", "Hello world.")
	f.SetCellValue("Blah", "B2", 100)

	// Set active sheet of the workbook
	f.SetActiveSheet(index)

	// Save xlsx file by the given path
	fileName := "GoBook.xlsx"
	os.Remove(fileName)
	if err := f.SaveAs(fileName); err != nil {
		println(err.Error())
	}
}

package main

import (
	//_ "image/gif"
	//_ "image/jpeg"
	_ "image/png"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	f := excelize.NewFile()

	// Insert a picture.
	if err := f.AddPicture("Sheet1", "A2", "mypng.png", ""); err != nil {
		println(err.Error())
	}

	// Insert a picture to worksheet with scaling.
	if err := f.AddPicture("Sheet1", "D2", "mypng.png", `{"x_scale": 0.5, "y_scale": 0.5}`); err != nil {
		println(err.Error())
	}
	// Insert a picture offset in the cell with printing support.
	if err := f.AddPicture("Sheet1", "H2", "mypng.png", `{"x_offset": 15, "y_offset": 10, "print_obj": true, "lock_aspect_ratio": false, "locked": false}`); err != nil {
		println(err.Error())
	}

	// Save the xlsx file with the origin path.
	if err := f.SaveAs("Pictures.xlsx"); err != nil {
		println(err.Error())
	}
}

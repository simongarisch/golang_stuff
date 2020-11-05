// go run dataframego.go
package main

import (
	"context"
	"fmt"
	"strings"

	dataframe "github.com/rocketlaunchr/dataframe-go"
	"github.com/rocketlaunchr/dataframe-go/imports"
)

func main() {
	s1 := dataframe.NewSeriesInt64("day", nil, 1, 2, 3, 4, 5, 6, 7, 8)
	s2 := dataframe.NewSeriesFloat64("sales", nil, 50.3, 23.4, 56.2, nil, nil, 84.2, 72, 89)
	df := dataframe.NewDataFrame(s1, s2)

	fmt.Print(df.Table())

	// insert and remove
	df.Append(nil, 9, 123.6)

	df.Append(nil, map[string]interface{}{
		"day":   10,
		"sales": nil,
	})

	df.Remove(0)
	fmt.Print(df.Table())

	// reading from csv
	ctx := context.TODO()

	csvStr := `
	Country,Date,Age,Amount,Id
	US,2012-02-01,50,112.1,01234
	UK,2012-02-01,32,321.31,54320
	UK,2012-02-01,17,18.2,12345
	US,2012-02-01,32,321.31,54320
	UK,2012-05-07,NA,18.2,12345
	US,2012-02-01,32,321.31,54320
	US,2012-02-01,32,321.31,54320
	Spain,2012-02-01,66,555.42,00241`
	df, err := imports.LoadFromCSV(ctx, strings.NewReader(csvStr))
	if err != nil {
		panic(err)
	}
	fmt.Print(df.Table())
}

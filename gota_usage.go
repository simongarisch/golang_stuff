// go run gota_usage.go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

func main() {
	csvfile, err := os.Open("test.csv")
	if err != nil {
		log.Fatal(err)
	}

	df := dataframe.ReadCSV(csvfile)
	fmt.Println(df)

	df = df.Filter(dataframe.F{"3", "==", 3})
	fmt.Println(df)

	s := series.New([]string{"b", "a", "c", "d"}, series.String, "SeriesName")
	fmt.Println(s)

	s1 := series.New([]string{"b", "a"}, series.String, "COL.1")
	s2 := series.New([]int{1, 2}, series.Int, "COL.2")
	s3 := series.New([]float64{3.0, 4.0}, series.Float, "COL.3")
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)

	dfx := dataframe.New(
		series.New([]string{"b", "a"}, series.String, "COL.1"),
		series.New([]int{1, 2}, series.Int, "COL.2"),
		series.New([]float64{3.0, 4.0}, series.Float, "COL.3"),
	)
	fmt.Println(dfx)

	fmt.Println("***********************")
	fmt.Println("Or load data from structs...")
	type User struct {
		Name     string
		Age      int
		Accuracy float64
		ignored  bool
	}
	users := []User{
		{"Aram", 17, 0.2, true},
		{"Juan", 18, 0.8, true},
		{"Ana", 22, 0.5, true},
	}
	dfs := dataframe.LoadStructs(users)
	fmt.Println(dfs)
	dfs = dfs.Filter(dataframe.F{"Age", ">=", 18})
	fmt.Println(dfs)
	fmt.Println(dfs.Select([]string{"Name", "Accuracy"}))

	// see more at https://github.com/go-gota/gota
}

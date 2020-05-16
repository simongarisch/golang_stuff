/*
The Separation of Concerns (SoC) is a design principle for separating a computer program into distinct sections.
With the Model View Controller (MVC) pattern we separate the out the data model, our visual representation or view,
and the controller that decides what model data to show on the view.
*/

package main

import (
	"fmt"
	"strings"
)

type PriceModel struct {
	stocks map[string]string
}

func NewPriceModel() {
	model := PriceModel{}
	model.stocks = map[string]string{
		"FB":   "Facebook",
		"AMZN": "Amazon",
		"AAPL": "Apple",
		"NFLX": "Netflix",
		"GOOG": "Google",
	}
}

func (model PriceModel) GetStockName(stockCode string) (string, error) {
	name, ok := model.stocks[stockCode]
	if !ok {
		return name, fmt.Errorf("stock code %q not recognised", stockCode)
	}
	return name, nil
}

func (model PriceModel) GetStockCode(stockName string) (string, error) {
	for code, name := range model.stocks {
		if name == stockName {
			return code, nil
		}
	}
	return "", fmt.Errorf("stock name %q not recognised", stockName)
}

type PriceView struct{}

func NewPriceView() {
	fmt.Println("***** MVC Example *****")
	details := `
		"Here we can do lookups for FAANG stocks! Options are:"
		"- 's' to lookup by stock name"
		"- 't' to lookup by stock ticker"
		"- 'e' to exit"
	`
	fmt.Println(details)
	underline()
}

func underline() {
	fmt.Println(strings.Repeat("*", 10))
}

func main() {
	NewPriceView()
}

package main

import "fmt"

func main() {

	stocks := map[string]float64{
		"AMZN": 1699.8,
		"GOOG": 1129.129,
		"MSFT": 98.61,
	}

	fmt.Println(len(stocks))    // 3
	fmt.Println(stocks["MSFT"]) // 98.61
	fmt.Println(stocks["TSLA"]) // 0, so zero value if not found
	value, ok := stocks["TSLA"]
	if !ok {
		fmt.Println("TSLA not found")
	} else {
		fmt.Println(value)
	}

	// now set TSLA
	stocks["TSLA"] = 322.12
	fmt.Println(stocks)

	// and delete a different stock
	delete(stocks, "AMZN")

	// now loop
	for key := range stocks {
		fmt.Println(key)
	}

	for key, value := range stocks {
		fmt.Println(key, value)
	}

}

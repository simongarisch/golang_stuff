package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("ini")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	fmt.Println(viper.Get("app.global1"))       // someValue
	fmt.Println(viper.Get("emails.general"))    // general@gmail.com
	fmt.Println(viper.Get("employees.manager")) // dave
}

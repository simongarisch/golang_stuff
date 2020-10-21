// go run viper.go
package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var (
	defaults = map[string]interface{}{
		"username": "admin",
		"password": "mypass",
		"host":     "localhost",
		"port":     3306,
		"database": "test",
	}

	configName = "config"
)

func main() {
	for k, v := range defaults {
		viper.SetDefault(k, v)
	}

	viper.SetConfigName(configName)
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("could not read config file: %v", err)
	}

	fmt.Println("Username from viper: %s\n", viper.GetString("username"))
}

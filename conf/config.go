package config

import (
	"log"

	"github.com/spf13/viper"
	// "gopkg.in/validator.v2"
)

func init() {
	readConfig()
	// registerValidator()
}

func readConfig() {
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetConfigName("app.config")

	err := viper.ReadInConfig()

	if err != nil {
		log.Println(err.Error())
	}
}

// func registerValidator() {
// 	validator.SetValidationFunc("notzz", lib.NotZZ)
// }

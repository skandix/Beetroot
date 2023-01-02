package config

import (
	"log"

	"github.com/spf13/viper"
)


func SetupConfig() {
	viper.SetConfigName("config")          // name of config file (without extension)
	viper.SetConfigType("yaml")            // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("$HOME/.beetroot") // call multiple times to add many search pathsw
	viper.AddConfigPath(".")               // optionally look for config in the working directory

	if err := viper.ReadInConfig(); err != nil { // Handle errors reading the config file
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			log.Fatalf("Error cannot find config file %s\n", err.Error())
		}
	}
}

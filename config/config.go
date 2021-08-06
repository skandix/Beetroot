package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func SetDefaults() {

}

func main() {
	viper.SetConfigName("config")          // name of config file (without extension)
	viper.SetConfigType("yaml")            // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("$HOME/.beetroot") // call multiple times to add many search pathsw
	viper.AddConfigPath(".")               // optionally look for config in the working directory

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
}

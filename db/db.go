package db

import (
	"fmt"

	"github.com/betauia/beetroot/config"
)

func DBConnect() {
	fmt.Println("hello from db pkg")
	config.SetupConfig()
}

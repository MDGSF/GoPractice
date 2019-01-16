package main

import (
	"fmt"

	"github.com/MDGSF/GoPractice/cobraDemo/cmd"
	"github.com/spf13/viper"
)

func main() {
	fmt.Println("test")
	cmd.Execute()
	viper.SetConfigName("test")
}

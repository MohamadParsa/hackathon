package main

import (
	"fmt"

	"github.com/MohamadParsa/hackathon/internal/adapter/in/restFull"
	"github.com/spf13/viper"
)

func main() {
	configViper()
	restFullServer := restFull.New()
	restFullServer.Serve(viper.GetString("serverPort"))
}
func configViper() {
	viper.AutomaticEnv()
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("failed to read config file: %v", err))
	}
}

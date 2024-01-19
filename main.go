package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/MohamadParsa/hackathon/internal/adapter/in/restFull"
	"github.com/MohamadParsa/hackathon/internal/adapter/out/db"
	"github.com/MohamadParsa/hackathon/internal/application/quickAccess"
	"github.com/MohamadParsa/hackathon/internal/application/suggestion"
	"github.com/spf13/viper"
)

func main() {
	configViper()
	dbAdapter, err := db.New("postgresql://root:SgAyaKgLgNzOpW4GDh4TnixU@makalu.liara.cloud:30005/postgres")
	if err != nil {
		log.Error("error in db", err)
	}
	quickAccessApplication := quickAccess.New(dbAdapter)
	suggestionApplication := suggestion.New()

	restFullServer := restFull.New(quickAccessApplication, suggestionApplication)
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

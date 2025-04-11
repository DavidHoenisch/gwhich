package config

import (
	//"github.com/BurntSushi/toml"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/DavidHoenisch/gwhich/models/config"
	"github.com/DavidHoenisch/gwhich/utils"
)

func parseConfigFile() (toml.MetaData, error) {
	configFile, err := utils.GetUserConfigPath()
	if err != nil {
		log.Println(err)
	}

	var apps []config.Application

	return toml.DecodeFile(configFile, apps)
}

func NewConfig() []config.Application {

	return nil
}

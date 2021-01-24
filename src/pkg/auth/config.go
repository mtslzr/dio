package auth

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
)

const (
	configDir       = "$HOME/.config/dio"
	configFile      = "config.yml"
	msgNoConfigFile = "No token currently configured. Add one with: `dio auth -t <token>`"
)

func checkConfig() {
	cfg, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", configDir, configFile))
	if cfg == nil {
		log.Infof("No configuration file found.")
		fmt.Println(msgNoConfigFile)
		return
	}
	if err != nil {
		log.Errorf("Error reading config file: %+v", err)
		return
	}
}

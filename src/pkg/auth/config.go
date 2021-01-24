package auth

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/user"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

const (
	configDir         = ".config/dio"
	configFile        = "token.yml"
	configTemplateDir = "src/pkg/auth" // TODO: Fix this to work with standalone binary.

	msgConfigDelete = "Local configuration and token deleted."
	msgNoConfigFile = "No token currently configured."
)

type ConfigYaml struct {
	Github struct {
		User  string `json:"user"`
		Token string `json:"token"`
	} `json:"github.com"`
}

func checkConfig() error {
	homeDir, err := getHomeDir()
	if err != nil {
		return err
	}

	cfg, err := ioutil.ReadFile(fmt.Sprintf("%s/%s/%s", homeDir, configDir, configFile))
	if cfg == nil {
		log.Infof("No configuration file found.")
		fmt.Println(msgNoConfigFile)
		return nil
	}
	if err != nil {
		log.Errorf("Error reading config file: %+v", err)
		return err
	}

	fmt.Println(string(cfg))
	return nil
}

func createConfig(token string, user string) error {
	homeDir, err := getHomeDir()
	if err != nil {
		return err
	}

	err = os.MkdirAll(fmt.Sprintf("%s/%s", homeDir, configDir), 0755)
	if err != nil {
		log.Errorf("Error creating config directory: %+v", err)
		return err
	}

	data, err := parseYaml()
	if err != nil {
		return err
	}

	data.Github.User = user
	data.Github.Token = token
	dataYaml, err := yaml.Marshal(data)
	if err != nil {
		log.Errorf("Error marshaling YAML: %+v", err)
		return err
	}

	err = ioutil.WriteFile(fmt.Sprintf("%s/%s/%s", homeDir, configDir, configFile), dataYaml, 0600)
	if err != nil {
		log.Errorf("Error writing config file: %+v", err)
	}

	return err
}

func destroyConfig() error {
	homeDir, err := getHomeDir()
	if err != nil {
		return err
	}

	err = os.RemoveAll(fmt.Sprintf("%s/%s", homeDir, configDir))
	if err != nil {
		log.Errorf("Error deleting config folder: %+v", err)
		return err
	}

	fmt.Println(msgConfigDelete)
	return nil
}

func getHomeDir() (string, error) {
	usr, err := user.Current()
	if err != nil {
		log.Errorf("Error getting current home directory: %+v", err)
	}

	return usr.HomeDir, err
}

func parseYaml() (ConfigYaml, error) {
	var data ConfigYaml

	ymlFile, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", configTemplateDir, configFile))
	if err != nil {
		log.Errorf("Error reading config template: %+v", err)
		return data, err
	}

	err = yaml.Unmarshal(ymlFile, &data)
	if err != nil {
		log.Errorf("Error unmarshaling YAML: %+v", err)
	}

	return data, err
}

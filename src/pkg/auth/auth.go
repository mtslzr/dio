package auth

import (
	log "github.com/sirupsen/logrus"
)

// Authenticate writes a new personal token to config.
func Authenticate(token string, user string) error {
	log.Info("Setting up local config file...")
	return createConfig(token, user)
}

// Destroy removes existing token configuration.
func Destroy() error {
	log.Info("Removing existing token configuration...")
	return destroyConfig()
}

// Status gets the current status of token configuration.
func Status() error {
	log.Info("Checking current token configuration...")
	return checkConfig()
}
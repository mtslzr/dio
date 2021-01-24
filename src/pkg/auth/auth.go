package auth

import (
	log "github.com/sirupsen/logrus"
)

// Authenticate writes a new personal token to config.
func Authenticate(token string) {
	log.Info("Setting up local config file...")
}

// Destroy removes existing token configuration.
func Destroy() {
	log.Info("Removing existing token configuration...")
}

// Status gets the current status of token configuration.
func Status() {
	log.Info("Checking current token configuration...")
	checkConfig()
}
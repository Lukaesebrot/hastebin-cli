package config

import (
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

// Initialize initializes the current configuration file
func Initialize() error {
	// Define and create the configuration file path
	home, err := homedir.Dir()
	if err != nil {
		return err
	}
	configPath := filepath.Join(home, ".hastebincli")
	err = os.MkdirAll(configPath, os.ModePerm)
	if err != nil {
		return err
	}

	// Define the configuration file data
	viper.AddConfigPath(configPath)
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	// Set the defaults of the configuration
	viper.SetDefault("instance", "https://hasteb.in")
	viper.SetDefault("autoClip", true)

	// Safe-save the default configuration if it does not exist
	viper.SafeWriteConfig()

	// Read the configuration file
	err = viper.ReadInConfig()
	return err
}

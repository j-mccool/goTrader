package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config structure to hold the configuration data
type Config struct {
	AlpacaAPIKey    string `json:"apiKey"`
	AlpacaAPISecret string `json:"apiSecret"`
	AlpacaBaseUrl   string `json:"apiEndpoint"`
	Port            string `json:"port"`
}

// LoadConfig loads the configuration from a JSON file
func LoadConfig() (c Config, err error) {
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	readerr := viper.ReadInConfig()

	if readerr != nil {
		fmt.Println("Error reading in config", readerr)
	}
	if c.Port = viper.GetString("Port"); err != nil {
		return
	}
	if c.AlpacaAPIKey = viper.GetString("apiKey"); err != nil {
		return
	}
	if c.AlpacaAPISecret = viper.GetString("apiSecret"); err != nil {
		return
	}
	if c.AlpacaBaseUrl = viper.GetString("apiEndpoint"); err != nil {
		return
	}
	return c, nil
}
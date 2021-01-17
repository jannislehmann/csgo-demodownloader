package utils

import (
	"encoding/json"
	"log"
	"os"
)

var config Config

// Config holds the application configuration
type Config struct {
	HistoryAPIKey   string `json:"matchHistoryAuthenticationCode"`
	SteamAPIKey     string `json:"steamApiKey"`
	KnownMatchCode  string `json:"knownMatchCode"`
	SteamID         string `json:"steamId"` // should be uint64
	Username        string `json:"username"`
	Password        string `json:"password"`
	TwoFactorSecret string `json:"twoFactorSecret"`
	DemosDir        string `json:"demosDir"`
}

// GetConfiguration returns the Config information
func GetConfiguration() Config {
	if config.DemosDir == "" {
		file := "./configs/config.json"
		configFile, err := os.Open(file)
		if err != nil {
			configFile.Close()
			log.Fatal(err)
		}
		jsonParser := json.NewDecoder(configFile)
		newConfig := Config{}
		err = jsonParser.Decode(&newConfig)
		if err != nil {
			log.Fatal(err)
		}
		defer configFile.Close()

		config = newConfig
	}

	return config
}

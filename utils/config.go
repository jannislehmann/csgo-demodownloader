package utils

import (
	"encoding/json"
	"log"
	"os"
)

var config Config

// Config holds the application configuration
type Config struct {
	DemosDir string        `json:"demosDir"`
	Steam    *SteamConfig  `json:"steam"`
	CSGO     []*CSGOConfig `json:"csgo"`
}

// SteamConfig holds the configuration about the steam account to use for communicating with the GameCoordinator.
type SteamConfig struct {
	SteamAPIKey     string `json:"apiKey"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	TwoFactorSecret string `json:"twoFactorSecret"`
}

// CSGOConfig holds the accounts to watch.
type CSGOConfig struct {
	HistoryAPIKey  string `json:"matchHistoryAuthenticationCode"`
	KnownMatchCode string `json:"knownMatchCode"`
	SteamID        string `json:"steamId"` // should be uint64
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

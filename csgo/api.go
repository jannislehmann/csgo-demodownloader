package csgo

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"

	"github.com/Cludch/csgo-demodownloader/utils"
)

// MatchResponse contains information about the latest match
type MatchResponse struct {
	Result struct {
		Nextcode string `json:"nextcode"`
	} `json:"result"`
}

// GetLatestMatch returns the latest match's share code.
func GetLatestMatch() string {
	config := utils.GetConfiguration()

	// Get latest match
	u, err := url.Parse("https://api.steampowered.com/ICSGOPlayers_730/GetNextMatchSharingCode/v1")
	if err != nil {
		log.Fatal(err)
	}

	// Build query
	q := u.Query()
	q.Set("key", config.Steam.SteamAPIKey)
	q.Set("steamid", config.CSGO[0].SteamID)
	q.Set("steamidkey", config.CSGO[0].HistoryAPIKey)
	q.Set("knowncode", config.CSGO[0].KnownMatchCode)
	u.RawQuery = q.Encode()

	// Request match code
	matchResponse := &MatchResponse{}

	r, err := http.Get(u.String())
	if err != nil {
		log.Fatal(err)
	}

	errJSON := json.NewDecoder(r.Body).Decode(matchResponse)

	if errJSON != nil {
		r.Body.Close()
		log.Fatal(err)
	}

	defer r.Body.Close()

	return matchResponse.Result.Nextcode
}

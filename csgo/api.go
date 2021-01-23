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

// GetNextMatch returns the next match's share code.
// It uses the saved share codes as the current one.
func GetNextMatch(csgoConfig *utils.CSGOConfig, steamAPIKey string) string {
	// Get latest match
	u, err := url.Parse("https://api.steampowered.com/ICSGOPlayers_730/GetNextMatchSharingCode/v1")
	if err != nil {
		log.Fatal(err)
	}

	// Get knowncode from the database
	latestShareCode := utils.GetLatestShareCode(csgoConfig.SteamID)

	// Build query
	q := u.Query()
	q.Set("key", steamAPIKey)
	q.Set("steamid", csgoConfig.SteamID)
	q.Set("steamidkey", csgoConfig.HistoryAPIKey)
	q.Set("knowncode", latestShareCode)
	u.RawQuery = q.Encode()

	matchResponse := &MatchResponse{}

	// Request match code
	r, err := http.Get(u.String())
	if err != nil {
		log.Print(err)
		return ""
	}

	// Forbidden = wrong api keys
	// Precondition Failed = Know match code or steam id wrong
	if r.StatusCode == http.StatusForbidden || r.StatusCode == http.StatusPreconditionFailed {
		r.Body.Close()
		csgoConfig.Disabled = true
		return ""
	}

	// Accepted means that there is no recent match code available
	if r.StatusCode == http.StatusAccepted {
		r.Body.Close()
		return ""
	}

	errJSON := json.NewDecoder(r.Body).Decode(matchResponse)

	if errJSON != nil {
		r.Body.Close()
		log.Print(err)
		return ""
	}

	defer r.Body.Close()

	return matchResponse.Result.Nextcode
}

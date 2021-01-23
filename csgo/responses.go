package csgo

import (
	"log"
	"time"

	"github.com/Cludch/csgo-demodownloader/utils"
)

// HandleGCReady starts a daemon and checks every hour for new demos.
func (c *CS) HandleGCReady(e *GCReadyEvent) {
	// Download all recents games from the logged in account
	c.GetRecentGames()

	// Request demos for share codes from the config
	// Add known shares codes from the config to the database.
	for _, csgoUser := range utils.GetConfiguration().CSGO {
		shareCode := csgoUser.KnownMatchCode
		c.RequestMatch(shareCode)
		utils.AddShareCode(csgoUser.SteamID, shareCode)
	}

	t := time.NewTicker(time.Minute)
	for {
		log.Println("checking for a new demo...")
		go c.GetDemos()
		<-t.C
	}
}

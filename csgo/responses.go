package csgo

import (
	"log"
	"time"
)

// HandleGCReady starts a daemon and checks every hour for new demos.
func (c *CS) HandleGCReady(e *GCReadyEvent) {
	// Download all recents games from the logged in account
	c.GetRecentGames()

	t := time.NewTicker(time.Hour)
	for {
		log.Println("checking for a new demo...")
		go c.GetDemos()
		<-t.C
	}
}

// HandleMatchDownloaded logs information about the downloaded demo.
func HandleMatchDownloaded(e *GCMatchDownloaded) {
	log.Printf("Downloaded demo %s\n", e.DemoName)
}

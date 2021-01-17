package csgo

import (
	"log"
	"strconv"

	"github.com/Cludch/csgo-demodownloader/csgo/protocol"
	"github.com/Cludch/csgo-demodownloader/utils"
	"github.com/Philipp15b/go-steam/protocol/gamecoordinator"
)

// HandleMatchList handles a gc message containing matches and tries to download those.
func (c *CS) HandleMatchList(packet *gamecoordinator.GCPacket) error {
	matchList := new(protocol.CMsgGCCStrike15V2_MatchList)
	packet.ReadProtoMsg(matchList)

	for _, match := range matchList.GetMatches() {
		for _, round := range match.GetRoundstatsall() {
			// Demo link is only linked in the last round and in this case the reserveration id is set.
			if round.GetReservationid() == 0 {
				continue
			}

			matchID := match.GetMatchid()
			demoname := utils.GetConfiguration().DemosDir + strconv.FormatUint(matchID, 10) + ".dem"
			url := round.GetMap()

			if utils.CheckIfMatchExistsAlready(matchID) {
				continue
			}

			err := utils.DownloadDemo(url, demoname)
			if err != nil {
				log.Fatal(err)
				continue
			}

			c.emit(&GCMatchDownloaded{DemoName: demoname, MatchID: matchID})
		}
	}

	return nil
}

// HandleClientWelcome creates a ready event and tries sends a command to download recent games.
func (c *CS) HandleClientWelcome(packet *gamecoordinator.GCPacket) error {
	log.Println("connected to csgo gc")
	if c.isConnected {
		return nil
	}

	c.isConnected = true

	c.emit(&GCReadyEvent{})

	return nil
}

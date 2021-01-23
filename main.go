package main

import (
	"log"

	"github.com/Philipp15b/go-steam"
	"github.com/Philipp15b/go-steam/protocol/steamlang"
	"github.com/Philipp15b/go-steam/totp"

	"github.com/Cludch/csgo-demodownloader/csgo"
	"github.com/Cludch/csgo-demodownloader/utils"
)

var config utils.Config
var csgoClient *csgo.CS

func main() {
	err := steam.InitializeSteamDirectory()

	if err != nil {
		log.Fatal(err)
	}

	config = utils.GetConfiguration()

	go utils.ScanDemosDir()

	totpInstance := totp.NewTotp(config.Steam.TwoFactorSecret)

	myLoginInfo := new(steam.LogOnDetails)
	myLoginInfo.Username = config.Steam.Username
	myLoginInfo.Password = config.Steam.Password
	twoFactorCode, err := totpInstance.GenerateCode()

	if err != nil {
		log.Fatal(err)
	}

	myLoginInfo.TwoFactorCode = twoFactorCode

	client := steam.NewClient()
	client.Connect()
	for event := range client.Events() {
		switch e := event.(type) {
		case *steam.ConnectedEvent:
			log.Print("connected to steam. Logging in...")
			client.Auth.LogOn(myLoginInfo)
		case *steam.LoggedOnEvent:
			log.Print("logged on")
			client.Social.SetPersonaState(steamlang.EPersonaState_Online)
			csgoClient = csgo.NewCSGO(client)
			csgoClient.SetPlaying(true)
			csgoClient.ShakeHands()
		case *csgo.GCReadyEvent:
			csgoClient.HandleGCReady(e)
		case *csgo.GCMatchDownloaded:
			csgo.HandleMatchDownloaded(e)
		case steam.FatalErrorEvent:
			log.Fatal(e)
		}
	}
}

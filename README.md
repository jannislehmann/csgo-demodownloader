# CSGO Demo Downloader

The CSGO demo downloader can automatically download new CSGO official matchmaking demos using the GameCoordinator.
In order to do this, a few API credentials and a **separate** Steam account is needed. The application creates a CSGO game sessions using the separate account
and uses the Steam web API to check whether a new demo can be fetched. If that is the case, the application sends a full match info request to the game's GameCoordinator.
The GC then returns information about the match which also contain a download link.

The tool saves all the match ids from the demos in the `demos` directory. This is used to prevent downloading a demo every hour.
At a later point, I plan to extend the toolset with a separate basic demo analyzer.

The sqlite database file is located in the `configs` directory. The database holds information about downloaded matches and will be extended in the future.

## Usage

Get the latest binary and set up your demo location and the config file.

### `config.json`

Copy the `config.json.example` in the `configs` dir and rename it to `config.json` in the same dir.

| Key   |      Value      |  Explanation |
|----------|-------------:|------:|
| `matchHistoryAuthenticationCode` |  `1234-ABCDE-5678`  | The match history authentication code can be generated [here](https://help.steampowered.com/en/wizard/HelpWithGameIssue/?appid=730&issueid=128) |
| `knownMatchCode` | `CSGO-abcde-efghi-jklmn-opqrs-tuvwx` |  A share code from one of your latest matches. Can be received via the Game -> Matches |
| `steamId` |   `e`   |  The SteamID64 of the account to watch |
| `steamApiKey` |   `d`   | The Steam Web API key. Can be generate [here](https://steamcommunity.com/dev/apikey) |
| `username` |   `c`   |  Steam username |
| `password` |   `b`   |  Steam password |
| `twoFactorSecret` |   `a`   | Base64 encoded two factor secret. Can be generated using e.g. the [Steam Desktop Authenticator](https://github.com/Jessecar96/SteamDesktopAuthenticator) |
| `demosDir` |   `/demos`   | The directory, where the demos should be stored. Should be an absolute path |

## Disclaimer

This is my first ever Golang project, thus you might find some bad practice and a few performance issues in the long run.
The project structure is horrible but it works for now. If you have suggestions please go a head and create an issue. This would help me a lot!

This tool is not affiliated with Valve Software or Steam.

## Other projects that helped me a lot

* [go-steam](https://github.com/Philipp15b/go-steam)
* [cs-go](https://github.com/Gacnt/cs-go)
* [go-dota2](https://github.com/paralin/go-dota2)
* [csgo](https://github.com/ValvePython/csgo)

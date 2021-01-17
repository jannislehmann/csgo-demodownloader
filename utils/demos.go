package utils

import (
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

// ScanDemosDir scans the demos dir and adds all match ids to the database
func ScanDemosDir() {
	err := filepath.Walk(GetConfiguration().DemosDir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if info.IsDir() {
				return nil
			}

			// Ignore non .dem-files
			if filepath.Ext(path) != ".dem" {
				return nil
			}

			filename := info.Name()
			demoName := strings.TrimSuffix(filename, filepath.Ext(filename))

			matchID := getMatchIDFromFileName(demoName)

			if matchID == 0 {
				return nil
			}

			AddMatchToDatabase(matchID)

			return nil
		})
	if err != nil {
		panic(err)
	}
}

// Filename without extension
func getMatchIDFromFileName(demoName string) uint64 {
	// Demos downloaded by this tool are ${matchid}.dem
	// Demos downloaded by the game are match730_${matchid}_${outcomeid}_${token}.dem
	// Downloaded demos via a share code through the game are ${matchid}_${outcomeid}.dem

	var matchIDString string

	demoNameParts := strings.Split(demoName, "_")

	// Check for match730_${matchid}_${outcomeid}_${token}.dem
	reGameOwn := regexp.MustCompile(`^match730(_?[\d]{21})(_?[\d]{10})(_?[\d]{3})$`)
	foundGameOwn := reGameOwn.MatchString(demoName)

	if foundGameOwn {
		matchIDString = demoNameParts[1]
	}

	// Check for ${matchid}_${outcomeid}.dem
	reGameShareCode := regexp.MustCompile(`^(_?[\d]{21})(_?[\d]{10})$`)
	foundGameShareCode := reGameShareCode.MatchString(demoName)

	// Check for ${matchid}.dem
	reTool := regexp.MustCompile(`^[0-9]*$`)
	foundTool := reTool.MatchString(demoName)

	if foundGameShareCode || foundTool {
		matchIDString = demoNameParts[0]
	}

	if matchIDString == "" {
		return 0
	}

	matchID, err := strconv.ParseUint(matchIDString, 10, 64)

	if err != nil {
		return 0
	}

	return matchID
}

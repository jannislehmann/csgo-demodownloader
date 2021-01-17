// Handles file download and extraction

package utils

import (
	"compress/bzip2"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
)

type invalidDownloadURLError struct{}

type demoNotFoundError struct{}

func (e *invalidDownloadURLError) Error() string {
	return "Invalid download url"
}

func (e *demoNotFoundError) Error() string {
	return "Demo no longer downloadable"
}

// DownloadDemo will download a demo from an url and decompress and store it in local filepath.
// It writes to the destination file as it downloads it, without
// loading the entire file into memory.
func DownloadDemo(url string, filepath string) error {
	// Validate the url
	re := regexp.MustCompile(`^http:\/\/replay[\d]{3}\.valve\.net\/730\/[\d]{21}_([\d]*)\.dem\.bz2$`)

	if !re.MatchString(url) {
		return &invalidDownloadURLError{}
	}

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		log.Println(err)
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url) //nolint // We have to take dynamic replay urls in order to download them. Url is validated before.
	if err != nil || resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return &demoNotFoundError{}
	}

	// Decompress and write to file
	cr := bzip2.NewReader(resp.Body)
	_, err = io.Copy(out, cr)

	defer resp.Body.Close()

	if err != nil {
		return err
	}

	log.Printf("Downloaded demo %s\n", filepath)

	return nil
}

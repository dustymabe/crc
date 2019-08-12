package download

import (
	"github.com/code-ready/crc/pkg/crc/logging"

	"github.com/cavaliercoder/grab"
	"github.com/code-ready/crc/pkg/crc/errors"
)

func Download(uri, destination string) (string, error) {
	// create client
	logging.Debugf("Downloading %s to %s", uri, destination)
	client := grab.NewClient()
	req, err := grab.NewRequest(destination, uri)
	if err != nil {
		return "", errors.Newf("Not able to get response from %s: %v", uri, err)
	}
	defer func() {
		if err != nil {
			os.Remove(destination)
		}
	}()
	resp := client.Do(req)
	// check for errors
	if err := resp.Err(); err != nil {
		return "", errors.Newf("Download failed: %v\n", err)
	}

	logging.Debugf("Download saved to %v \n", resp.Filename)
	return resp.Filename, nil
}

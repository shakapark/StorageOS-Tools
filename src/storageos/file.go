package storageos

import (
	"errors"
	"io/ioutil"
	"os"
)

// GetFileID Get NodeID in local file
func GetFileID(path string) (string, error) {

	dat, err := ioutil.ReadFile(path)
	if err != nil {
		return "", errors.New("Reading File: " + error.Error(err))
	}

	return string(dat), nil
}

// ReplaceFileID Replace ID in file
func ReplaceFileID(path, newID string) error {
	f, err := os.Create(path)
	if err != nil {
		return errors.New("Opening File: " + error.Error(err))
	}
	defer f.Close()

	_, err = f.WriteString(newID + "\n")
	if err != nil {
		return errors.New("Writing File: " + error.Error(err))
	}

	return nil
}

package storageos

import (
	"errors"
	"os"
)

// GetFileID Get NodeID in local file
func GetFileID(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", errors.New("Opening File: " + error.Error(err))
	}
	defer f.Close()

	b1 := make([]byte, 64)
	n1, err := f.Read(b1)
	if err != nil {
		return "", errors.New("Reading File: " + error.Error(err))
	}

	if b1[n1-1] == 10 {
		b1 = b1[:(n1 - 1)]
	}

	return string(b1), nil
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

package file

import (
	"encoding/json"
	"os"
)

func Write(pathToFile string, data interface{}) error {
	f, err := os.OpenFile(pathToFile, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if _, err := f.Write(bytes); err != nil {
		return err
	}

	return nil
}

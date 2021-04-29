package file

import (
	"io/ioutil"
	"os"
)

func Read(pathToFile string) ([]byte, error) {
	jsonFile, err := os.Open(pathToFile)

	if err != nil {
		return nil, err
	}

	defer jsonFile.Close()

	data, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	return data, nil
}

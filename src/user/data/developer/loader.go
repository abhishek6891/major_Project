package developer

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func LoadDevelopersFromFile(filePath string) ([]Developer, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var developers []Developer
	err = json.Unmarshal(byteValue, &developers)
	if err != nil {
		return nil, err
	}

	return developers, nil
}

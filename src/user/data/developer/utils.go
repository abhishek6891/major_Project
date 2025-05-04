package developer

import (
	"encoding/json"
	"io/ioutil"
)

const FilePath = "Doc/project/developer.json"

func LoadDevelopers() ([]Developer, error) {
	file, err := ioutil.ReadFile(FilePath)
	if err != nil {
		return nil, err
	}
	var developers []Developer
	err = json.Unmarshal(file, &developers)
	return developers, err
}

func SaveDevelopers(developers []Developer) error {
	data, err := json.MarshalIndent(developers, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(FilePath, data, 0644)
}

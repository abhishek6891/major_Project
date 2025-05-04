package projects

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func LoadProjectsFromFile(filePath string) ([]Project, error) {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}

	var projects []Project
	err = json.Unmarshal(byteValue, &projects)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json: %v", err)
	}

	return projects, nil
}

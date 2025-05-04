package projects

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

const filePath = "Doc/project/project.json"

func LoadProjects() ([]Project, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var projects []Project
	err = json.Unmarshal(data, &projects)
	return projects, err
}

func SaveProjects(projects []Project) error {
	data, err := json.MarshalIndent(projects, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filePath, data, os.ModePerm)
}

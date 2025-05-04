package client

import (
	"encoding/json"
	"io/ioutil"
)

const FilePath = "Doc/project/client.json"

func LoadClients() ([]Client, error) {
	file, err := ioutil.ReadFile(FilePath)
	if err != nil {
		return nil, err
	}
	var clients []Client
	err = json.Unmarshal(file, &clients)
	return clients, err
}

func SaveClients(projects []Client) error {
	data, err := json.MarshalIndent(projects, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(FilePath, data, 0644)
}

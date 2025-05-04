package client

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func LoadClientsFromFile(filePath string) ([]Client, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var clients []Client
	err = json.Unmarshal(byteValue, &clients)
	if err != nil {
		return nil, err
	}

	return clients, nil
}

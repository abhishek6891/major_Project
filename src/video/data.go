package video

import (
	"encoding/json"
	"os"
)

func SaveCallSession(session VideoCall, path string) error {
	calls := []VideoCall{}

	file, err := os.ReadFile(path)
	if err == nil {
		json.Unmarshal(file, &calls)
	}
	calls = append(calls, session)

	data, _ := json.MarshalIndent(calls, "", "  ")
	return os.WriteFile(path, data, 0644)
}

package chat

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func SaveChatSession(session ChatSession) error {
	dir := "Doc/project/chat_history"
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	// File name based on client, developer, and hour
	fileName := fmt.Sprintf("%s_%s_%s.json", "clientID"+session.ClientID, "DeveloperID"+session.DeveloperID, session.StartTime.Format("2006-01-02_15"))
	filePath := filepath.Join(dir, fileName)

	var existingSessions []ChatSession

	// If file exists, load it
	if data, err := os.ReadFile(filePath); err == nil {
		_ = json.Unmarshal(data, &existingSessions)
	}

	existingSessions = append(existingSessions, session)

	data, err := json.MarshalIndent(existingSessions, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filePath, data, 0644)
}

package chat

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func ChatHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	defer ws.Close()

	clientID := r.URL.Query().Get("client_id")
	developerID := r.URL.Query().Get("developer_id")
	projectName := r.URL.Query().Get("project_name")

	session := ChatSession{
		ClientID:    clientID,
		DeveloperID: developerID,
		ProjectName: projectName,
		StartTime:   time.Now(),
		Messages:    []ChatMessage{},
	}

	log.Printf("Chat started between client %s and developer %s", clientID, developerID)

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			log.Println("Connection closed:", err)
			break
		}

		message := ChatMessage{
			SenderID:    r.URL.Query().Get("sender_id"),
			SenderName:  r.URL.Query().Get("sender_name"),
			SenderType:  r.URL.Query().Get("sender_type"),
			ProjectName: projectName,
			Message:     string(msg),
			Time:        time.Now(),
		}

		session.Messages = append(session.Messages, message)

		if err := ws.WriteMessage(websocket.TextMessage, msg); err != nil {
			log.Println("Write error:", err)
			break
		}
	}

	session.EndTime = time.Now()
	session.Duration = session.EndTime.Sub(session.StartTime).String()

	if err := SaveChatSession(session); err != nil {
		log.Println("Error saving session:", err)
	} else {
		log.Println("Chat session saved.")
	}
}

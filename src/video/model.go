package video

import "time"

type VideoCall struct {
	CallerID  string    `json:"callerId"`
	CalleeID  string    `json:"calleeId"`
	RoomID    string    `json:"roomId"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime,omitempty"`
}

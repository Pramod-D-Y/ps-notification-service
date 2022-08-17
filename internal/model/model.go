package model

type Notification struct {
	
	Message     string             `json:"message,omitempty"`
	To          []string             `json:"to,omitempty"`
}

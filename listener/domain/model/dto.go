package model

type ResponseData struct {
	TxID      string `json:"txID"`
	Source    string `json:"source"`
	EventType string `json:"eventType"`
	Payload   PayloadData
}

type RequestData struct {
	TxID      string      `json:"txID"`
	Source    string      `json:"source"`
	EventType string      `json:"eventType"`
	Payload   PayloadData `json:"payload"`
}

type RequestEventCreate struct {
	EventType      string `json:"eventType"`
	EventName      string `json:"eventName"`
	Price          int    `json:"price"`
	MaxParticipant int    `json:"maxParticipant"`
}

type RequestEventUpdated struct {
	EventType string `json:"eventType"`
	ID        int    `json:"id"`
	EventName string `json:"eventName"`
	Price     int    `json:"price"`
}

type RequestParticipantEnrolled struct {
	EventType      string `json:"eventType"`
	ID             int    `json:"id"`
	ParticipantIDs []int  `json:"participantIds"`
}

type PayloadData struct {
	Message string `json:"message"`
	Email   string `json:"email"`
}

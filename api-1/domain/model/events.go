package model

type EventsInsertRequest struct {
	EventName      string `json:"eventName"`
	Price          int    `json:"price"`
	MaxParticipant int    `json:"maxParticipant"`
}

type EventsUpdateRequest struct {
	EventName string `json:"eventName"`
	Price     int    `json:"price"`
}

type EnrollEventsRequest struct {
	ParticipantIds []int `json:"participantIds"`
}

type EventsResponse struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Errors  error       `json:"errors"`
}

package entity

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type IntSlice []int

func (i *IntSlice) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to unmarshal JSONB value: %v", value)
	}
	return json.Unmarshal(bytes, i)
}

func (i IntSlice) Value() (driver.Value, error) {
	return json.Marshal(i)
}

type EventsModel struct {
	ID               int        `json:"id" gorm:"column:ID"`
	EventName        string     `json:"eventName" gorm:"column:event_name"`
	Price            int        `json:"price" gorm:"column:price"`
	ParticipantIds   IntSlice   `json:"participantIds" gorm:"participant_ids"` // This is not a column in the database, just for convenience
	MaxParticipant   int        `json:"maxParticipant" gorm:"column:max_participant"`
	TotalParticipant int        `json:"totalParticipant" gorm:"column:total_participant"`
	IsFullRegistered bool       `json:"isFullRegistered" gorm:"column:is_full_registered"`
	IsFullPaid       bool       `json:"isFullPaid" gorm:"column:is_full_paid"`
	CreatedDate      *time.Time `json:"createdDate" gorm:"column:created_date"`
	CreatedBy        *string    `json:"createdBy" gorm:"column:created_by"`
	UpdatedDate      *time.Time `json:"updatedDate" gorm:"column:updated_date"`
	UpdatedBy        *string    `json:"updatedBy" gorm:"column:updated_by"`
}

type ParticipantModel struct {
	ID       int          `json:"id" gorm:"column:id"` // userId
	EventsId int          `json:"-"`
	Events   *EventsModel `gorm:"foreignKey:EventsId" json:"events,omitempty"`
	Name     string       `json:"name" gorm:"column:name"`
	Paid     bool         `json:"paid" gorm:"column:paid"`
}

func (t EventsModel) TableName() string {
	return "Events-QCA"
}
func (t ParticipantModel) TableName() string {
	return "Participant-QCA"
}

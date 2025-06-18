package events

import (
	"go-starter-api/domain/entity"
	"go-starter-api/domain/model"

	"gorm.io/gorm"
)

type eventsRepository struct {
	db *gorm.DB
}

type EventsRepository interface {
	// GetCustomerAll(model []entity.CustomerModel) ([]entity.CustomerModel, error)
	InsertEvents(model model.EventsInsertRequest) (entity.EventsModel, error)
	UpdateEvents(model model.EventsUpdateRequest, id int) (entity.EventsModel, error)
}

func NewEventsRepository(db *gorm.DB) EventsRepository {
	return eventsRepository{db}
}

type EventsDto struct {
	FirstName string `json:"firstName" gorm:"column:FIRST_NAME"`
}

func (t EventsDto) TableName() string {
	return "Events-QCA"
}

// func (c eventsRepository) GetCustomerAll(model []entity.CustomerModel) ([]entity.CustomerModel, error) {
// 	tx := c.db.Find(&model).Limit(100)
// 	if tx.Error != nil {
// 		return nil, tx.Error
// 	}
// 	return model, nil
// }

func (c eventsRepository) InsertEvents(req model.EventsInsertRequest) (entity.EventsModel, error) {
	res := entity.EventsModel{
		EventName:        req.EventName,
		Price:            req.Price,
		MaxParticipant:   req.MaxParticipant,
		IsFullRegistered: false,
		IsFullPaid:       false,
	}

	tx := c.db.Omit("CreatedDate", "CreatedBy", "UpdatedDate", "UpdatedBy").Create(&res)
	if tx.Error != nil {
		return entity.EventsModel{}, tx.Error
	}
	return res, nil
}

func (c eventsRepository) UpdateEvents(req model.EventsUpdateRequest, id int) (entity.EventsModel, error) {
	res := entity.EventsModel{
		ID:        id,
		EventName: req.EventName,
		Price:     req.Price,
	}

	tx := c.db.Model(&res).Omit("CreatedDate", "CreatedBy", "UpdatedDate", "UpdatedBy").Where("ID = ?", id).Updates(&res)
	if tx.Error != nil {
		return entity.EventsModel{}, tx.Error
	}
	return res, nil
}

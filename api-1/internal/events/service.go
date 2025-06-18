package events

import (
	"go-starter-api/domain/entity"
	"go-starter-api/domain/model"
)

type eventsService struct {
	eventsRepository EventsRepository
}

type EventsService interface {
	// GetCustomerAll() ([]entity.CustomerModel, error)
	InsertEvents(req model.EventsInsertRequest) (entity.EventsModel, error)
	UpdateEvents(req model.EventsUpdateRequest, id int) (entity.EventsModel, error)
}

func NewEventsService(eventsRepository EventsRepository) EventsService {
	return eventsService{eventsRepository}
}

// func (c eventsService) GetCustomerAll() ([]entity.CustomerModel, error) {
// 	model := []entity.CustomerModel{}
// 	res, err := c.eventsRepository.GetCustomerAll(model)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return res, nil
// }

func (c eventsService) InsertEvents(req model.EventsInsertRequest) (entity.EventsModel, error) {
	res, err := c.eventsRepository.InsertEvents(req)
	if err != nil {
		return entity.EventsModel{}, err
	}
	return res, nil
}

func (c eventsService) UpdateEvents(req model.EventsUpdateRequest, id int) (entity.EventsModel, error) {
	res, err := c.eventsRepository.UpdateEvents(req, id)
	if err != nil {
		return entity.EventsModel{}, err
	}
	return res, nil
}

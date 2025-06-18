package events

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-starter-api/domain/model"
	"go-starter-api/pkg/utils"

	"gitlab.com/banpugroup/banpucoth/itsddev/library/golang/go-azure-sdk.git/appinsightsx"

	"github.com/gofiber/fiber/v2"
)

type eventsHandler struct {
	eventsService EventsService
	validate      *utils.Validator
	ai            appinsightsx.Appinsightsx
}

type EventsHandler interface {
	// GetCustomerAll(c *fiber.Ctx) error
	InsertEvents(c *fiber.Ctx) error
	UpdateEvents(c *fiber.Ctx) error
	EnrollEvents(c *fiber.Ctx) error
}

func NewEventsHandler(eventsService EventsService, validate *utils.Validator, ai appinsightsx.Appinsightsx) EventsHandler {
	return eventsHandler{eventsService, validate, ai}
}

// GetCustomerAll godoc
// @Summary Get all customers
// @Security BearerAuth
// @Description Get all customers
// @Tags Customer
// @Accept json
// @Produce json
// @Success 200 {object} model.CustomerResponse
// @Failure 400 {object} utils.ErrorDTO
// @Router /customer [get]

// func (cus eventsHandler) GetCustomerAll(c *fiber.Ctx) error {
// 	data, err := cus.customerService.GetCustomerAll()
// 	if err != nil {
// 		cus.ai.Error(appinsightsx.LoggerRequest{
// 			Error: err.Error(),
// 		})
// 		return c.Status(fiber.StatusBadRequest).JSON(err)
// 	}

// 	res := model.CustomerResponse{
// 		Data:    data,
// 		Message: "success",
// 		Errors:  nil,
// 	}

// 	cus.ai.Info(appinsightsx.LoggerRequest{
// 		Process: fmt.Sprintf("GetCustomerAll: %v", res),
// 	})
// 	return c.JSON(res)
// }

// InsertCustomer handles the insertion of a new customer.
// @Summary Insert a new customer
// @Security BearerAuth
// @Description Insert a new customer into the system
// @Tags Customer
// @Accept json
// @Produce json
// @Param customer body model.CustomerInsertRequest true "Customer Insert Request"
// @Success 200 {object} model.CustomerResponse
// @Failure 400 {object} utils.ErrorDTO
// @Router /customer [post]
func (cus eventsHandler) InsertEvents(c *fiber.Ctx) error {
	req := model.EventsInsertRequest{}
	err := c.BodyParser(&req)
	if err != nil {
		cus.ai.Error(appinsightsx.LoggerRequest{
			Error: err.Error(),
		})
		return c.Status(fiber.StatusBadRequest).JSON(utils.NewErrorDTO(fiber.StatusBadRequest, "", err))
	}

	v := cus.validate.Validate(req)
	if len(v) > 0 {
		jsonString, err := json.Marshal(v)
		if err != nil {
			cus.ai.Error(appinsightsx.LoggerRequest{
				Error: err.Error(),
			})
			return c.Status(fiber.StatusBadRequest).JSON(utils.NewErrorDTO(fiber.StatusBadRequest, "", err))
		}
		cus.ai.Error(appinsightsx.LoggerRequest{
			Error: fmt.Errorf("validate len 0"),
		})
		return c.Status(fiber.StatusBadRequest).JSON(utils.NewErrorDTO(fiber.StatusBadRequest, "", errors.New(string(jsonString))))
	}

	data, err := cus.eventsService.InsertEvents(req)
	if err != nil {
		cus.ai.Error(appinsightsx.LoggerRequest{
			Error: err.Error(),
		})
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	res := model.EventsResponse{
		Data:    data,
		Message: "success",
		Errors:  nil,
	}

	cus.ai.Info(appinsightsx.LoggerRequest{
		Process: fmt.Sprintf("InsertEvents: %v", res),
	})
	return c.JSON(res)
}

// InsertCustomer handles the insertion of a new customer.
// @Summary Insert a new customer
// @Security BearerAuth
// @Description Insert a new customer into the system
// @Tags Customer
// @Accept json
// @Produce json
// @Param customer body model.CustomerInsertRequest true "Customer Insert Request"
// @Success 200 {object} model.CustomerResponse
// @Failure 400 {object} utils.ErrorDTO
// @Router /customer [post]
func (cus eventsHandler) UpdateEvents(c *fiber.Ctx) error {
	req := model.EventsUpdateRequest{}
	err := c.BodyParser(&req)
	if err != nil {
		cus.ai.Error(appinsightsx.LoggerRequest{
			Error: err.Error(),
		})
		return c.Status(fiber.StatusBadRequest).JSON(utils.NewErrorDTO(fiber.StatusBadRequest, "", err))
	}

	v := cus.validate.Validate(req)
	if len(v) > 0 {
		jsonString, err := json.Marshal(v)
		if err != nil {
			cus.ai.Error(appinsightsx.LoggerRequest{
				Error: err.Error(),
			})
			return c.Status(fiber.StatusBadRequest).JSON(utils.NewErrorDTO(fiber.StatusBadRequest, "", err))
		}
		cus.ai.Error(appinsightsx.LoggerRequest{
			Error: fmt.Errorf("validate len 0"),
		})
		return c.Status(fiber.StatusBadRequest).JSON(utils.NewErrorDTO(fiber.StatusBadRequest, "", errors.New(string(jsonString))))
	}

	id, _ := c.ParamsInt("id")

	data, err := cus.eventsService.UpdateEvents(req, id)
	if err != nil {
		cus.ai.Error(appinsightsx.LoggerRequest{
			Error: err.Error(),
		})
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	res := model.CustomerResponse{
		Data:    data,
		Message: "success",
		Errors:  nil,
	}

	cus.ai.Info(appinsightsx.LoggerRequest{
		Process: fmt.Sprintf("InsertCustomer: %v", res),
	})
	return c.JSON(res)
}

func (cus eventsHandler) EnrollEvents(c *fiber.Ctx) error {
	// req := model.EnrollEventsRequest{}
	// fmt.Println(req)
	// err := c.BodyParser(&req)
	// if err != nil {
	// 	cus.ai.Error(appinsightsx.LoggerRequest{
	// 		Error: err.Error(),
	// 	})
	// 	return c.Status(fiber.StatusBadRequest).JSON(utils.NewErrorDTO(fiber.StatusBadRequest, "", err))
	// }

	// v := cus.validate.Validate(req)
	// if len(v) > 0 {
	// 	jsonString, err := json.Marshal(v)
	// 	if err != nil {
	// 		cus.ai.Error(appinsightsx.LoggerRequest{
	// 			Error: err.Error(),
	// 		})
	// 		return c.Status(fiber.StatusBadRequest).JSON(utils.NewErrorDTO(fiber.StatusBadRequest, "", err))
	// 	}
	// 	cus.ai.Error(appinsightsx.LoggerRequest{
	// 		Error: fmt.Errorf("validate len 0"),
	// 	})
	// 	return c.Status(fiber.StatusBadRequest).JSON(utils.NewErrorDTO(fiber.StatusBadRequest, "", errors.New(string(jsonString))))
	// }

	// id, _ := c.ParamsInt("id")

	// data, err := cus.eventsService.UpdateEvents(req, id)
	// if err != nil {
	// 	cus.ai.Error(appinsightsx.LoggerRequest{
	// 		Error: err.Error(),
	// 	})
	// 	return c.Status(fiber.StatusBadRequest).JSON(err)
	// }

	// res := model.CustomerResponse{
	// 	Data:    data,
	// 	Message: "success",
	// 	Errors:  nil,
	// }

	// cus.ai.Info(appinsightsx.LoggerRequest{
	// 	Process: fmt.Sprintf("InsertCustomer: %v", res),
	// })
	// return c.JSON(res)
	return nil
}

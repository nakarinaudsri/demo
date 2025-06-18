package events

import (
	"go-starter-api/pkg/utils"

	"gitlab.com/banpugroup/banpucoth/itsddev/library/golang/go-azure-sdk.git/appinsightsx"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Router(r fiber.Router, db *gorm.DB, ai appinsightsx.Appinsightsx) {
	validate := utils.NewValidator()

	repository := NewEventsRepository(db)
	service := NewEventsService(repository)
	handler := NewEventsHandler(service, validate, ai)

	eventsRoute := r.Group("/events")
	// groupRoute.Get("", handler.GetCustomerAll)
	eventsRoute.Post("/", handler.InsertEvents)
	eventsRoute.Patch("/:id", handler.UpdateEvents)
	eventsRoute.Post("/:id/enroll", handler.EnrollEvents)
}

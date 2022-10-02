package service

import (
	"github.com/onemgvv/WB_L2/develop/dev11/internal/delivery/http/dto"
	event "github.com/onemgvv/WB_L2/develop/dev11/internal/domain"
	"github.com/onemgvv/WB_L2/develop/dev11/pkg/storage"
)

type Calendar interface {
	CreateEvent(dto dto.CreateEventDTO) string
	ReadEvent(uid string) (*event.Event, error)
	UpdateEvent(uid string, dto dto.UpdateEventDTO) error
	DeleteEvent(uid string) error
}

type Deps struct {
	*storage.Storage
}

type Service struct {
	Calendar
}

func NewService(deps Deps) *Service {
	return &Service{NewCalendarService(deps.Storage)}
}

package service

import (
	"github.com/onemgvv/WB_L2/develop/dev11/internal/delivery/http/dto"
	event "github.com/onemgvv/WB_L2/develop/dev11/internal/domain"
	"github.com/onemgvv/WB_L2/develop/dev11/pkg/storage"
)

type CalendarService struct {
	*storage.Storage
}

func NewCalendarService(storage *storage.Storage) *CalendarService {
	return &CalendarService{Storage: storage}
}

func (c *CalendarService) CreateEvent(dto dto.CreateEventDTO) string {
	ev := dto.ToModel()
	c.Set(ev.Id, *ev)
	return ev.Id
}

func (c *CalendarService) ReadEvent(uid string) (*event.Event, error) {
	ev, err := c.Get(uid)
	if err != nil {
		return &event.Event{}, err
	}

	return ev, nil
}

func (c *CalendarService) UpdateEvent(uid string, dto dto.UpdateEventDTO) error {
	return c.Update(uid, *dto.ToModel())
}

func (c *CalendarService) DeleteEvent(uid string) error {
	return c.Delete(uid)
}

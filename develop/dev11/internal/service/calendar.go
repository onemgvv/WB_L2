package service

import (
	"fmt"
	"time"

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

func (c *CalendarService) DailyEvents(userId int, date time.Time) event.Events {
	var events event.Events

	for _, event := range c.GetByUser(userId) {
		if event.Date.Day() == date.Day() {
			events = append(events, event)
		}
	}

	return events
}

func (c *CalendarService) WeeklyEvents(userId int, date time.Time) event.Events {
	start := date.Add(-3 * time.Hour)
	end := date.AddDate(0, 0, 7).Add(3 * time.Hour)
	var events event.Events

	fmt.Println(end)

	for _, event := range c.GetByUser(userId) {
		if event.Date.After(start) && event.Date.Before(end) {
			events = append(events, event)
		}
	}

	return events
}

func (c *CalendarService) MonthlyEvents(userId int, date time.Time) event.Events {
	start := date.Add(-3 * time.Hour)
	end := date.AddDate(0, 1, 0).Add(3 * time.Hour)
	var events event.Events

	for _, event := range c.GetByUser(userId) {
		if event.Date.After(start) && event.Date.Before(end) {
			events = append(events, event)
		}
	}

	return events
}
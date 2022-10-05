package domain

import (
	"time"

	"github.com/onemgvv/WB_L2/develop/dev11/pkg/uuid"
)

type Events []Event

type Event struct {
	Id          string    `json:"id"`
	Title       string    `json:"name"`
	Description string    `json:"description"`
	UserId      int       `json:"userId"`
	Date        time.Time `json:"date"`
}

func NewEvent(title, description string, date time.Time, userId int) *Event {
	uid := uuid.V4()
	if date.IsZero() {
		date = time.Now()
	}

	return &Event{uid, title, description, userId, date}
}

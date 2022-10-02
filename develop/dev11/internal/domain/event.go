package domain

import (
	"github.com/onemgvv/WB_L2/develop/dev11/pkg/uuid"
	"time"
)

type Event struct {
	Id          string    `json:"id"`
	Title       string    `json:"name"`
	Description string    `json:"description"`
	UserId      int       `json:"userId"`
	Date        time.Time `json:"date"`
}

func NewEvent(title, description string, userId int) *Event {
	uid := uuid.V4()
	now := time.Now()

	return &Event{uid, title, description, userId, now}
}

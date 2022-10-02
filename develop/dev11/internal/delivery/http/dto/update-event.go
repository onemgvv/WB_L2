package dto

import (
	"github.com/onemgvv/WB_L2/develop/dev11/internal/domain"
	"time"
)

type UpdateEventDTO struct {
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	UserId      int       `json:"userId,omitempty"`
	Date        time.Time `json:"date,omitempty"`
}

func (u *UpdateEventDTO) ToModel() *domain.Event {
	return domain.NewEvent(u.Title, u.Description, u.UserId)
}

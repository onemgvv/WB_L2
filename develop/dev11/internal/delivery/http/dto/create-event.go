package dto

import (
	"github.com/onemgvv/WB_L2/develop/dev11/internal/domain"
	"time"
)

type CreateEventDTO struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	UserId      int       `json:"userId"`
	Date        time.Time `json:"date"`
}

func (u *CreateEventDTO) ToModel() *domain.Event {
	return domain.NewEvent(u.Title, u.Description, u.UserId)
}

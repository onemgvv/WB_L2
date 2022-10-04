package dto

import (
	"log"
	"strconv"
	"time"

	"github.com/onemgvv/WB_L2/develop/dev11/internal/domain"
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

func ParseUpdateEvDto(title, userId, description, date string) *UpdateEventDTO {
	uId, err := strconv.Atoi(userId)
	if err != nil {
		log.Printf("%s.\n", err.Error())
		return nil
	}

	cTime, err := time.Parse("2006-01-02", date)
	if err != nil {
		log.Printf("%s.\n", err.Error())
		return nil
	}
	
	return &UpdateEventDTO{
		Title: title,
		Description: description,
		UserId: uId,
		Date: cTime,
	}
}